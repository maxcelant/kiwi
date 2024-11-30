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
					tokens := []lexer.Token{
						{
							Type:    lexer.NIL,
							Literal: "nil",
							Lexeme:  "nil",
							Line:    1,
						},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: nil}))
				})
			})

			When("its a list of just one true token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{
						{
							Type:    lexer.TRUE,
							Literal: "true",
							Lexeme:  "true",
							Line:    1,
						},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: true}))
				})
			})

			When("its a list of just one false token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{
						{
							Type:    lexer.FALSE,
							Literal: "false",
							Lexeme:  "false",
							Line:    1,
						},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: false}))
				})
			})

			When("its a list of just one string token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{
						{
							Type:    lexer.STRING,
							Literal: "foo",
							Lexeme:  "\"foo\"",
							Line:    1,
						},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: "foo"}))
				})
			})

			When("its a list of just one number token", func() {
				It("returns a tree with just one primary type node", func() {
					tokens := []lexer.Token{
						{
							Type:    lexer.NUMBER,
							Literal: 5,
							Lexeme:  "5",
							Line:    1,
						},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Primary{value: 5}))
				})
			})
		})

		Describe("Unary", func() {
			When("its a list with a negative and number", func() {
				It("returns a tree with a unary node", func() {
					tokens := []lexer.Token{
						{Type: lexer.MINUS, Lexeme: "-", Line: 1},
						{Type: lexer.NUMBER, Literal: 5, Lexeme: "5", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Unary{
						operator: tokens[0],
						right:    &Primary{value: 5},
					}))
				})
			})

			When("its a list with a bang and number", func() {
				It("returns a tree with a unary node", func() {
					tokens := []lexer.Token{
						{Type: lexer.BANG, Lexeme: "!", Line: 1},
						{Type: lexer.NUMBER, Literal: 5, Lexeme: "5", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Unary{
						operator: tokens[0],
						right:    &Primary{value: 5},
					}))
				})
			})
		})

		Describe("Factor", func() {
			When("its a list with two numbers and an a slash", func() {
				It("returns a tree with one factor node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.SLASH, Lexeme: "/", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left:     &Primary{value: 1},
						operator: tokens[1],
						right:    &Primary{value: 2},
					}))
				})
			})

			When("its a list with two numbers and an a star", func() {
				It("returns a tree with one factor node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.STAR, Lexeme: "*", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left:     &Primary{value: 1},
						operator: tokens[1],
						right:    &Primary{value: 2},
					}))
				})
			})

			When("its a list of multiple numbers and star tokens", func() {
				It("returns a nested factor tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.STAR, Lexeme: "*", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.STAR, Lexeme: "*", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left: &Binary{
							left:     &Primary{value: 1},
							operator: tokens[1],
							right:    &Primary{value: 1},
						},
						operator: tokens[3],
						right:    &Primary{value: 1},
					}))
				})
			})
		})

		Describe("Term", func() {
			When("its a list with two numbers and a plus", func() {
				It("returns a tree with one term node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.PLUS, Lexeme: "+", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left:     &Primary{value: 1},
						operator: tokens[1],
						right:    &Primary{value: 2},
					}))
				})
			})

			When("its a list with two numbers and a minus", func() {
				It("returns a tree with one term node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.MINUS, Lexeme: "-", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left:     &Primary{value: 1},
						operator: tokens[1],
						right:    &Primary{value: 2},
					}))
				})
			})

			When("its a list of multiple numbers and plus tokens", func() {
				It("returns a nested term tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.PLUS, Lexeme: "+", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.PLUS, Lexeme: "+", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual).To(Equal(&Binary{
						left: &Binary{
							left:     &Primary{value: 1},
							operator: tokens[1],
							right:    &Primary{value: 1},
						},
						operator: tokens[3],
						right:    &Primary{value: 1},
					}))
				})
			})
		})
	})
})
