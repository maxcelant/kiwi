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
				actual, err := parser.Parse()
				Expect(err).To(BeNil())
				Expect(actual).To(BeNil())
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
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: nil}))
				})
			})

			When("its a list of just one true token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{{
						Type:    lexer.TRUE,
						Literal: "true",
						Lexeme:  "true",
						Line:    1,
					}}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: true}))
				})
			})

			When("its a list of just one false token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{{
						Type:    lexer.FALSE,
						Literal: "false",
						Lexeme:  "false",
						Line:    1,
					}}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: false}))
				})
			})

			When("its a list of just one string token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{{
						Type:    lexer.STRING,
						Literal: "foo",
						Lexeme:  "\"foo\"",
						Line:    1,
					}}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: "foo"}))
				})
			})

			When("its a list of just one number token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{{
						Type:    lexer.NUMBER,
						Literal: 5,
						Lexeme:  "5",
						Line:    1,
					}}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: 5}))
				})
			})
		})
	})
})
