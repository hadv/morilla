package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMorilla(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Morilla Suite")
}
