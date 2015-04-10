package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/rightscale/rsc.v1-unstable" // import "gopkg.in/rightscale/rsc.v1-unstable"
)

var _ = Describe("Encrypt", func() {
	Context("given a string value", func() {
		var (
			seekret = "sensitive value"
		)

		It("encrypts", func() {
			var _, err = main.Encrypt(seekret)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("decrypts", func() {
			var encrypted, err = main.Encrypt(seekret)
			Ω(err).ShouldNot(HaveOccurred())
			decrypted, err := main.Decrypt(encrypted)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(decrypted).Should(Equal(seekret))
		})

	})

})
