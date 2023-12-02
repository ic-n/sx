package contracts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestContracts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contracts Suite")
}
