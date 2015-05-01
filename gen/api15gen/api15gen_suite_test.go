package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApi15gen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api15gen Suite")
}
