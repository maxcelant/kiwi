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

	Context("when sent an empty string", func() {
		It("should return an empty list of tokens", func() {
			in := ""
			result, _ := lexer.ScanLine(in)
			Expect(result).To(Equal([]Token{}))
		})
	})

	Context("when sent a single char identifier", func() {
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

		When("there is a single and multi-string version", func() {
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

	Context("when sent a multi-string identifier", func() {
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
})
