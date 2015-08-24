package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAPI15gen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API15gen Suite")
}
