package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCmds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmds Suite")
}
