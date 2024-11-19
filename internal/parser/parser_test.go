package parser

import (
	"testing"

	"github.com/maxcelant/kiwi/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Parser", func() {

	Describe("Empty", func() {
		When("its an empty list of tokens", func() {
			It("returns an empty parse tree", func() {
				tokens := []lexer.Token{}
				parser := New(tokens)
				actual, err := parser.parse()
				Expect(err).To(BeNil())
				Expect(actual).To(Equal())
			})
		})
	})

	Describe("Expressions", func() {
		Describe("Primary", func() {
			When("its a list of just one nil token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{{
						Type:    lexer.NIL,
						Literal: "nil",
						Lexeme:  "nil",
						Line:    1,
					}}
					parser := New(tokens)
					actual, err := parser.parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(Primary{value: "nil"}))
				})
			})
		})
	})
})
