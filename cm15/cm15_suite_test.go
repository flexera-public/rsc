package cm15_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCm15(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cm15 Suite")
}
