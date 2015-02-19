package main_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc"
)

var _ = Describe("Config", func() {
	Context("loading a config", func() {
		var (
			path   string
			config *main.ClientConfig
			err    error
		)

		JustBeforeEach(func() {
			config, err = main.LoadConfig(path)
		})

		Context("with an incorrect path", func() {
			BeforeEach(func() {
				path = "/does/not/exist"
			})

			It("returns an error", func() {
				Ω(err).Should(HaveOccurred())
				Ω(config).Should(BeNil())
			})
		})

		Context("with a file", func() {
			var (
				tempFile *os.File
			)

			BeforeEach(func() {
				tempFile, _ = ioutil.TempFile("", "rsc_test")
				path = tempFile.Name()
			})

			AfterEach(func() {
				os.Remove(tempFile.Name())
			})

			Context("containing invalid JSON", func() {
				BeforeEach(func() {
					tempFile.WriteString(`{invalidJson}`)
				})

				It("returns an error", func() {
					Ω(err).Should(HaveOccurred())
					Ω(config).Should(BeNil())
				})
			})

			Context("containing valid JSON but not a config", func() {
				BeforeEach(func() {
					tempFile.WriteString(`{"a":1}`)
				})

				It("loads an empty config", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(config).ShouldNot(BeNil())
					Ω(config.Account).Should(BeZero())
					Ω(config.Host).Should(BeZero())
					Ω(config.Token).Should(BeZero())
				})
			})

			Context("containing a config", func() {
				var (
					account int
					host    string
					token   string
				)

				JustBeforeEach(func() {
					var cfg = fmt.Sprintf(`{"Account":%d,"Host":"%s","Token":"%s"}`,
						account, host, token)
					tempFile.WriteString(cfg)
					config, err = main.LoadConfig(path)
				})

				Context("which is valid", func() {
					var (
						tok = "seekret"
					)

					BeforeEach(func() {
						account = 42
						host = "Host"
						token, _ = main.Encrypt(tok)
					})

					It("loads", func() {
						Ω(err).ShouldNot(HaveOccurred())
						Ω(config).ShouldNot(BeNil())
						Ω(config.Account).Should(Equal(account))
						Ω(config.Host).Should(Equal(host))
						Ω(config.Token).Should(Equal(tok))
					})

				})

				Context("with an non encrypted token", func() {
					BeforeEach(func() {
						account = 42
						host = "host"
						token = "tok"
					})

					It("returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})
				})
			})

		})

	})

	Context("Creating a config", func() {
		var (
			path    string
			err     error
			testOut *bytes.Buffer
			testIn  io.Reader
		)

		JustBeforeEach(func() {
			main.SetOutput(testOut)
			main.SetInput(testIn)
			err = main.CreateConfig(path)
		})

		AfterEach(func() {
			main.SetOutput(os.Stdout)
			main.SetInput(os.Stdin)
		})

		Context("with valid input values", func() {
			BeforeEach(func() {
				testOut = &bytes.Buffer{}
				testIn = bytes.NewReader([]byte("\n\n\n\n"))
			})

			Context("with an invalid path", func() {
				BeforeEach(func() {
					path = "/does/not/exist"
				})

				It("returns an error", func() {
					Ω(err).Should(HaveOccurred())
				})
			})

			Context("with a valid path", func() {
				var (
					tempFile *os.File
				)

				BeforeEach(func() {
					tempFile, _ = ioutil.TempFile("", "rsc_test")
					path = tempFile.Name()
				})

				AfterEach(func() {
					os.Remove(tempFile.Name())
				})

				Context("which contains a valid config", func() {
					BeforeEach(func() {
						var cfg = `{"Account":2,"Host":"s","Token":"OlVr2Xv9jZfg1zf+LACM+WJNnFxg4Bm46Yc/kA=="}`
						tempFile.WriteString(cfg)
					})

					It("asks for confirmation", func() {
						Ω(testOut.String()).Should(ContainSubstring("overwrite?"))
					})

					It("exits without prompting for new values", func() {
						Ω(testOut.String()).Should(ContainSubstring("Exiting"))
						Ω(testOut.String()).ShouldNot(ContainSubstring("Account id"))
					})

					Context("replying yes to the confirmation prompt", func() {
						BeforeEach(func() {
							testIn = bytes.NewReader([]byte("y\n"))
						})

						It("prompts for new values", func() {
							Ω(testOut.String()).ShouldNot(ContainSubstring("Exiting"))
							Ω(testOut.String()).Should(ContainSubstring("Account id"))
						})
					})
				})

				Context("which contains an invalid config", func() {
					BeforeEach(func() {
						var cfg = `invalid`
						tempFile.WriteString(cfg)
					})

					It("does not ask for confirmation", func() {
						Ω(testOut.String()).ShouldNot(ContainSubstring("overwrite?"))
					})

					Context("setting inputs", func() {
						var (
							account string
							token   string
							host    string
						)

						Context("given an incorrect account", func() {
							BeforeEach(func() {
								account = "foo"
								var inputs = fmt.Sprintf("%s\n%s\n%s\n",
									account, token, host)
								testIn = bytes.NewReader([]byte(inputs))
							})

							It("returns a descriptive error", func() {
								Ω(err).Should(HaveOccurred())
								Ω(err.Error()).Should(ContainSubstring("Account ID must be an integer"))
							})
						})

						Context("given correct values", func() {
							BeforeEach(func() {
								account = "71"
								token = "tok"
								host = "host"
								var inputs = fmt.Sprintf("%s\n%s\n%s\n",
									account, token, host)
								testIn = bytes.NewReader([]byte(inputs))
							})

							It("saves the config", func() {
								Ω(err).ShouldNot(HaveOccurred())
								config, err := main.LoadConfig(tempFile.Name())
								Ω(err).ShouldNot(HaveOccurred())
								Ω(config.Account).Should(Equal(71))
								Ω(config.Token).Should(Equal(token))
								Ω(config.Host).Should(Equal(host))
							})
						})
					})
				})

			})

		})

	})

})
