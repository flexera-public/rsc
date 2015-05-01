package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encrypt", func() {
	Context("given a string value", func() {
		var (
			seekret = "sensitive value"
		)

		It("encrypts", func() {
			var _, err = Encrypt(seekret)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("decrypts", func() {
			var encrypted, err = Encrypt(seekret)
			Ω(err).ShouldNot(HaveOccurred())
			decrypted, err := Decrypt(encrypted)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(decrypted).Should(Equal(seekret))
		})

	})

})
