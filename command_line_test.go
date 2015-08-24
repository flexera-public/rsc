package main

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/cm16"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/ss"
	"gopkg.in/alecthomas/kingpin.v2"
)

var _ = Describe("Command line parsing", func() {

	Context("with a command line", func() {
		var (
			args []string
			app  *kingpin.Application
			err  error

			cmdLine *cmd.CommandLine
		)

		BeforeEach(func() {
			app = kingpin.New("test", "test")
		})

		JustBeforeEach(func() {
			os.Args = append([]string{"rsc"}, args...)
			cmdLine, err = ParseCommandLine(app)
		})

		Context("using cm15", func() {
			BeforeEach(func() {
				args = []string{"--noAuth", "--host=h", "cm15", "index", "clouds"}
			})

			Context("with x1", func() {
				BeforeEach(func() {
					args = append([]string{"--x1=foo"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.ExtractOneSelect).ShouldNot(BeEmpty())
				})
			})

			Context("with xm", func() {
				BeforeEach(func() {
					args = append([]string{"--xm=foo"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.ExtractSelector).ShouldNot(BeEmpty())
				})
			})

			Context("with xj", func() {
				BeforeEach(func() {
					args = append([]string{"--xj=foo"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.ExtractSelectorJSON).ShouldNot(BeEmpty())
				})
			})

			Context("with xh", func() {
				BeforeEach(func() {
					args = append([]string{"--xh=foo"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.ExtractHeader).ShouldNot(BeEmpty())
				})
			})

			Context("with fetch", func() {
				BeforeEach(func() {
					args = append([]string{"--fetch"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.FetchResource).Should(BeTrue())
				})
			})

			Context("with dump", func() {
				BeforeEach(func() {
					args = append([]string{"--dump=debug"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.Dump).Should(Equal("debug"))
				})
			})

			Context("with verbose", func() {
				BeforeEach(func() {
					args = append([]string{"--verbose"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.Verbose).Should(BeTrue())
				})
			})

			Context("with pp", func() {
				BeforeEach(func() {
					args = append([]string{"--pp"}, args...)
				})

				It("initializes the command line struct", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cmdLine.Pretty).Should(BeTrue())
				})
			})
		})

		Context("creating a client", func() {
			var (
				client cmd.CommandClient
				clErr  error
			)

			Context("using cm15", func() {
				BeforeEach(func() {
					args = []string{"--noAuth", "--host=h", "cm15", "index", "clouds"}
				})

				JustBeforeEach(func() {
					client, clErr = APIClient("cm15", cmdLine)
				})
				It("creates the client", func() {
					Ω(clErr).ShouldNot(HaveOccurred())
					Ω(client).ShouldNot(BeNil())
					_, ok := client.(*cm15.API)
					Ω(ok).Should(BeTrue())
				})
			})

			Context("using cm16", func() {
				BeforeEach(func() {
					args = []string{"--noAuth", "--host=h", "cm16", "index", "clouds"}
				})

				JustBeforeEach(func() {
					client, clErr = APIClient("cm16", cmdLine)
				})
				It("creates the client", func() {
					Ω(clErr).ShouldNot(HaveOccurred())
					Ω(client).ShouldNot(BeNil())
					_, ok := client.(*cm16.API)
					Ω(ok).Should(BeTrue())
				})
			})

			Context("using ss", func() {
				BeforeEach(func() {
					args = []string{"--noAuth", "--host=h", "ss", "index", "clouds"}
				})

				JustBeforeEach(func() {
					client, clErr = APIClient("ss", cmdLine)
				})
				It("creates the client", func() {
					Ω(clErr).ShouldNot(HaveOccurred())
					Ω(client).ShouldNot(BeNil())
					_, ok := client.(*ss.API)
					Ω(ok).Should(BeTrue())
				})
			})

		})

	})
})
