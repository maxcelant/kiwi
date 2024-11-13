package lexer

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

var _ = Describe("Lexer", func() {
	var lexer Lexer

	BeforeEach(func() {
		lexer = Lexer{}
	})

	Context("sent an empty char", func() {
		It("should return an empty list of tokens", func() {
			in := ""
			result, _ := lexer.ScanLine(in)
			Expect(result).To(Equal([]Token{}))
		})
	})

	Context("sent whitespace chars", func() {
		When("its a single whitespace", func() {
			It("should return an empty list of tokens", func() {
				in := " "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
			})
		})
		When("its a multiple whitespaces", func() {
			It("should return an empty list of tokens", func() {
				in := " "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
			})
		})
	})

	Context("sent a single char identifier", func() {
		It("should return a list with just a single token", func() {
			in := ";"
			result, _ := lexer.ScanLine(in)
			expected := Token{
				Type:    SEMICOLON,
				Literal: ";",
				Lexeme:  ";",
				Line:    1,
			}
			Expect(result).To(Equal([]Token{expected}))
		})

		When("there is a single and multi-char version", func() {
			It("should return a list with just a single token", func() {
				in := "="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    EQUAL,
					Literal: "=",
					Lexeme:  "=",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})
	})

	Context("sent a multi-char identifier", func() {
		When("its a two char token", func() {
			It("should return a list with just a single token", func() {
				in := "=="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    EQUAL_EQUAL,
					Literal: "==",
					Lexeme:  "==",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})
		When("its a 3+ char token", func() {
			It("should return a list with just a single string token", func() {
				in := "\"foobar\""
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STRING,
					Literal: "foobar",
					Lexeme:  "\"foobar\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})
	})
})
