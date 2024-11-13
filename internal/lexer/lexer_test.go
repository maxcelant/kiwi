package lexer

import (
	"errors"
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
	var eofToken Token

	BeforeEach(func() {
		lexer = Lexer{}
		eofToken = Token{Type: EOF, Lexeme: "", Literal: nil, Line: 1}
	})

	Context("empty", func() {
		It("should return an empty list of tokens", func() {
			in := ""
			result, _ := lexer.ScanLine(in)
			Expect(result).To(Equal([]Token{eofToken}))
		})
	})

	Context("whitespaces", func() {
		When("its a single whitespace", func() {
			It("should return an empty list of tokens", func() {
				in := " "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{eofToken}))
			})
		})
		When("its a multiple whitespaces", func() {
			It("should return an empty list of tokens", func() {
				in := " "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{eofToken}))
			})
		})
	})

	Context("single char identifier", func() {
		It("should return a list with just a single token", func() {
			in := ";"
			result, _ := lexer.ScanLine(in)
			expected := Token{
				Type:    SEMICOLON,
				Literal: ";",
				Lexeme:  ";",
				Line:    1,
			}
			Expect(result).To(Equal([]Token{expected, eofToken}))
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
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})
	})

	Context("strings", func() {
		When("its a multi-line char token", func() {
			It("should return a list with just a single string token", func() {
				in := "\"foobar\""
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STRING,
					Literal: "foobar",
					Lexeme:  "\"foobar\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})
	})

	Context("numbers", func() {
		When("its a single digit char token", func() {
			It("should return a list with just a single digit token", func() {
				in := "1"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    NUMBER,
					Literal: 1,
					Lexeme:  "1",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})

		When("its a multi-digit char token", func() {
			It("should return a list with just a single digit token", func() {
				in := "123"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    NUMBER,
					Literal: 123,
					Lexeme:  "123",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})

		When("its a digit with alpha characters in it", func() {
			It("should return an error", func() {
				in := "123abc"
				_, err := lexer.ScanLine(in)
				Expect(err).To(Equal(errors.New("invalid number: contains alphabetic characters")))
			})
		})
	})

	Context("comparators", func() {
		When("its an equal-equal sign", func() {
			It("should return a list with just a equal-equal token", func() {
				in := "=="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    EQUAL_EQUAL,
					Literal: "==",
					Lexeme:  "==",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})

		When("its a not-equal sign", func() {
			It("should return a list with just a not-equal token", func() {
				in := "!="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    BANG_EQ,
					Literal: "!=",
					Lexeme:  "!=",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected, eofToken}))
			})
		})
	})

	Context("multiple tokens", func() {
		When("given two single char tokens", func() {
			It("should return a list with two tokens", func() {
				in := "=;"
				result, _ := lexer.ScanLine(in)
				exp1 := Token{
					Type:    EQUAL,
					Literal: "=",
					Lexeme:  "=",
					Line:    1,
				}
				exp2 := Token{
					Type:    SEMICOLON,
					Literal: ";",
					Lexeme:  ";",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{exp1, exp2, eofToken}))
			})
		})

		When("given two string tokens", func() {
			It("should return a list with two string tokens", func() {
				in := "\"foo\" \"bar\""
				result, _ := lexer.ScanLine(in)
				exp1 := Token{
					Type:    STRING,
					Literal: "foo",
					Lexeme:  "\"foo\"",
					Line:    1,
				}
				exp2 := Token{
					Type:    STRING,
					Literal: "bar",
					Lexeme:  "\"bar\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{exp1, exp2, eofToken}))
			})
		})
	})
})
