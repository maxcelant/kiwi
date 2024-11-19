package parser

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Parser", func() {
	Describe("Primary", func() {
		When("its a list of just one nil token", func() {
			It("returns a tree with just one primary type node", func() {

			})
		})
	})
})
