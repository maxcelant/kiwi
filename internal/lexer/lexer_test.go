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

	BeforeEach(func() {
		lexer = Lexer{}
	})

	Context("empty", func() {
		It("should return an empty list of tokens", func() {
			in := ""
			result, _ := lexer.ScanLine(in)
			Expect(result).To(Equal([]Token{}))
		})
	})

	Context("whitespaces", func() {
		When("its a single whitespace", func() {
			It("should return an empty list of tokens", func() {
				in := " "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
			})
		})
		When("its a multiple whitespaces", func() {
			It("should return an empty list of tokens", func() {
				in := "   "
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
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

	Context("comments", func() {
		When("entering double slashes", func() {
			It("should return an empty list", func() {
				in := "// this is a comment"
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
			})
		})
		When("entering double slashes with newline", func() {
			It("should return an empty list", func() {
				in := "// this is a comment\n"
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{}))
			})
		})
	})

	Context("strings", func() {
		When("its a string with just alpha characters", func() {
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

		When("its a string with just numeric characters", func() {
			It("should return a list with just a single string token", func() {
				in := "\"123\""
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STRING,
					Literal: "123",
					Lexeme:  "\"123\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("its a string with alphanumeric characters", func() {
			It("should return a list with just a single string token", func() {
				in := "\"foobar123\""
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STRING,
					Literal: "foobar123",
					Lexeme:  "\"foobar123\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("its a string with a space in the middle", func() {
			It("should return a list with just a single string token", func() {
				in := "\"foo bar\""
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STRING,
					Literal: "foo bar",
					Lexeme:  "\"foo bar\"",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
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
				Expect(result).To(Equal([]Token{expected}))
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
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("its a multi-digit char token attached to a non-digit", func() {
			It("should return a list with two tokens", func() {
				in := "123;"
				result, _ := lexer.ScanLine(in)
				token1 := Token{
					Type:    NUMBER,
					Literal: 123,
					Lexeme:  "123",
					Line:    1,
				}
				token2 := Token{
					Type:    SEMICOLON,
					Literal: ";",
					Lexeme:  ";",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token1, token2}))
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
		When("its equal related", func() {
			It("should return a list with just a equal-equal token", func() {
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

			It("should return a list with just a not-equal token", func() {
				in := "!="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    BANG_EQ,
					Literal: "!=",
					Lexeme:  "!=",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("its less-than related", func() {
			It("should return a list with just a less-than token", func() {
				in := "<"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    LESS,
					Literal: "<",
					Lexeme:  "<",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})

			It("should return a list with just a less-than-or-equal token", func() {
				in := "<="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    LESS_EQ,
					Literal: "<=",
					Lexeme:  "<=",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("its greater-than related", func() {
			It("should return a list with just a greater-than token", func() {
				in := ">"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    GREATER,
					Literal: ">",
					Lexeme:  ">",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})

			It("should return a list with just a greater-than-or-equal token", func() {
				in := ">="
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    GREATER_EQ,
					Literal: ">=",
					Lexeme:  ">=",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})
	})

	Context("braces, parenthesis", func() {
		When("given a single left brace", func() {
			It("should return a list with just a single left brace token", func() {
				in := "{"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    LEFT_BRACE,
					Literal: "{",
					Lexeme:  "{",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("given a single left parenthesis", func() {
			It("should return a list with just a single left parenthesis token", func() {
				in := "("
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    LEFT_PAREN,
					Literal: "(",
					Lexeme:  "(",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})
	})

	Context("math symbols", func() {
		When("entering a slash", func() {
			It("should return a list with a div token", func() {
				in := "/"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    SLASH,
					Literal: "/",
					Lexeme:  "/",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("entering a plus", func() {
			It("should return a list with a plus token", func() {
				in := "+"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    PLUS,
					Literal: "+",
					Lexeme:  "+",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("entering a minus", func() {
			It("should return a list with a minus token", func() {
				in := "-"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    MINUS,
					Literal: "-",
					Lexeme:  "-",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("entering a star", func() {
			It("should return a list with a star token", func() {
				in := "*"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    STAR,
					Literal: "*",
					Lexeme:  "*",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
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
				Expect(result).To(Equal([]Token{exp1, exp2}))
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
				Expect(result).To(Equal([]Token{exp1, exp2}))
			})
		})
	})

	Context("keywords", func() {
		When("given a single keyword", func() {
			It("returns a list with a single keyword", func() {
				in := "var"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    VAR,
					Literal: "var",
					Lexeme:  "var",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("given a multiple keyword", func() {
			It("returns a list with a keywords", func() {
				in := "true false"
				result, _ := lexer.ScanLine(in)
				token1 := Token{
					Type:    TRUE,
					Literal: "true",
					Lexeme:  "true",
					Line:    1,
				}
				token2 := Token{
					Type:    FALSE,
					Literal: "false",
					Lexeme:  "false",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token1, token2}))
			})
		})

		When("given a the func keyword", func() {
			It("returns a token list with func keyword", func() {
				in := "fn"
				result, _ := lexer.ScanLine(in)
				token := Token{
					Type:    FUNC,
					Literal: "fn",
					Lexeme:  "fn",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token}))
			})
		})

		When("given the and keyword", func() {
			It("returns a token list with the and keyword", func() {
				in := "and"
				result, _ := lexer.ScanLine(in)
				token := Token{
					Type:    AND,
					Literal: "and",
					Lexeme:  "and",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token}))
			})
		})

		When("given the or keyword", func() {
			It("returns a token list with the or keyword", func() {
				in := "or"
				result, _ := lexer.ScanLine(in)
				token := Token{
					Type:    OR,
					Literal: "or",
					Lexeme:  "or",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token}))
			})
		})
		When("given the print keyword", func() {
			It("returns a token list with the print keyword", func() {
				in := "print"
				result, _ := lexer.ScanLine(in)
				token := Token{
					Type:    PRINT,
					Literal: "print",
					Lexeme:  "print",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token}))
			})
		})
	})

	Context("identifiers", func() {
		When("given a 3 letter identifier", func() {
			It("returns a list with a identifier token", func() {
				in := "foo"
				result, _ := lexer.ScanLine(in)
				expected := Token{
					Type:    IDENTIFIER,
					Literal: "foo",
					Lexeme:  "foo",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{expected}))
			})
		})

		When("given a var keyword and an identifier", func() {
			It("returns a list with both tokens", func() {
				in := "var foo"
				result, _ := lexer.ScanLine(in)
				token1 := Token{
					Type:    VAR,
					Literal: "var",
					Lexeme:  "var",
					Line:    1,
				}
				token2 := Token{
					Type:    IDENTIFIER,
					Literal: "foo",
					Lexeme:  "foo",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token1, token2}))
			})
		})

		When("given a full variable declaration", func() {
			It("returns a list with both tokens", func() {
				in := "var foo = 3"
				result, _ := lexer.ScanLine(in)
				Expect(result).To(Equal([]Token{
					{
						Type:    VAR,
						Literal: "var",
						Lexeme:  "var",
						Line:    1,
					},
					{
						Type:    IDENTIFIER,
						Literal: "foo",
						Lexeme:  "foo",
						Line:    1,
					},
					{
						Type:    EQUAL,
						Literal: "=",
						Lexeme:  "=",
						Line:    1,
					},
					{
						Type:    NUMBER,
						Literal: 3,
						Lexeme:  "3",
						Line:    1,
					},
				}))
			})
		})

		When("its an identifier with a parenthesis attached", func() {
			It("should return a list with just a identifier and a parenthesis", func() {
				in := "foo("
				result, _ := lexer.ScanLine(in)
				token1 := Token{
					Type:    IDENTIFIER,
					Literal: "foo",
					Lexeme:  "foo",
					Line:    1,
				}
				token2 := Token{
					Type:    LEFT_PAREN,
					Literal: "(",
					Lexeme:  "(",
					Line:    1,
				}
				Expect(result).To(Equal([]Token{token1, token2}))
			})
		})
	})
})
