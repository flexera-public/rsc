package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("loading a config", func() {
		var (
			path   string
			config *ClientConfig
			err    error
		)

		JustBeforeEach(func() {
			config, err = LoadConfig(path)
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
					tempFile.WriteString(`{invalidJSON}`)
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
					Ω(config.LoginHost).Should(BeZero())
					Ω(config.Email).Should(BeZero())
					Ω(config.Password).Should(BeZero())
					Ω(config.RefreshToken).Should(BeZero())
				})
			})

			Context("containing a config", func() {
				var (
					account  int
					host     string
					email    string
					password string
					refresh  string
				)

				JustBeforeEach(func() {
					var cfg = fmt.Sprintf(`{"Account":%d,"Email":"%s","LoginHost":"%s","Password":"%s","RefreshToken":"%s"}`,
						account, email, host, password, refresh)
					tempFile.WriteString(cfg)
					config, err = LoadConfig(path)
				})

				Context("which is valid", func() {
					var (
						tok = "seekret"
					)

					BeforeEach(func() {
						account = 42
						email = "test@test.com"
						host = "LoginHost"
						password, _ = Encrypt(tok)
					})

					It("loads", func() {
						Ω(err).ShouldNot(HaveOccurred())
						Ω(config).ShouldNot(BeNil())
						Ω(config.Account).Should(Equal(account))
						Ω(config.Email).Should(Equal(email))
						Ω(config.LoginHost).Should(Equal(host))
						Ω(config.Password).Should(Equal(tok))
					})

				})

				Context("with an non encrypted password", func() {
					BeforeEach(func() {
						account = 42
						host = "host"
						password = "tok"
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
			SetOutput(testOut)
			SetInput(testIn)
			err = CreateConfig(path)
		})

		AfterEach(func() {
			SetOutput(os.Stdout)
			SetInput(os.Stdin)
		})

		Context("with valid input values", func() {
			BeforeEach(func() {
				testOut = &bytes.Buffer{}
				testIn = bytes.NewReader([]byte("\n\n\n\n\n"))
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
						var cfg = `{"Account":2,"Email":"test@test.com","LoginHost":"s","Password":"OlVr2Xv9jZfg1zf+LACM+WJNnFxg4Bm46Yc/kA==","RefreshToken":"P+YuzH1milj3Od0Vd1tbQPAKWrVqXpRbTTtdeBc2U4HdH/tL2LqXEqp9OhDtWXB5slRWlZSjTVpjDjp/kpY9GJQX8D77nrY1"}`
						tempFile.WriteString(cfg)
					})

					It("asks for confirmation", func() {
						Ω(testOut.String()).Should(ContainSubstring("overwrite?"))
					})

					It("exits without prompting for new values", func() {
						Ω(testOut.String()).Should(ContainSubstring("Exiting"))
						Ω(testOut.String()).ShouldNot(ContainSubstring("Account ID"))
					})

					Context("replying yes to the confirmation prompt", func() {
						BeforeEach(func() {
							testIn = bytes.NewReader([]byte("y\n"))
						})

						It("prompts for new values", func() {
							Ω(testOut.String()).ShouldNot(ContainSubstring("Exiting"))
							Ω(testOut.String()).Should(ContainSubstring("Account ID"))
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
							account  string
							email    string
							password string
							host     string
							refresh  string
						)

						Context("given an incorrect account", func() {
							BeforeEach(func() {
								account = "foo"
								var inputs = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
									account, email, password, host, refresh)
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
								email = "test@test.com"
								password = "tok"
								host = "host"
								refresh = "1ecea8a3fbce2eb3a3be6ac8180f60182b5c81cd"
								var inputs = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
									account, email, password, host, refresh)
								testIn = bytes.NewReader([]byte(inputs))
							})

							It("saves the config", func() {
								Ω(err).ShouldNot(HaveOccurred())
								config, err := LoadConfig(tempFile.Name())
								Ω(err).ShouldNot(HaveOccurred())
								Ω(config.Account).Should(Equal(71))
								Ω(config.Email).Should(Equal(email))
								Ω(config.Password).Should(Equal(password))
								Ω(config.LoginHost).Should(Equal(host))
								Ω(config.RefreshToken).Should(Equal(refresh))
							})
						})
					})
				})

			})

		})

	})

})
