package rsapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRsapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rsapi Suite")
}
