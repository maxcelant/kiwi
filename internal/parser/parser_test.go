package parser

import (
	"testing"

	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
	"github.com/maxcelant/kiwi/internal/stmt"
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
				Expect(actual).To(BeEmpty())
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Primary{Value: nil}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Primary{Value: true}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Primary{Value: false}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Primary{Value: "foo"}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Primary{Value: 5}}))
				})
			})

			When("its a list with a number inside parentheses", func() {
				It("returns a tree with a grouping node", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_PAREN, Lexeme: "(", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.RIGHT_PAREN, Lexeme: ")", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Grouping{
						Expression: expr.Primary{Value: 1},
					}}))
				})
			})

			When("its a list with a number and a missing closing parenthesis", func() {
				It("returns an error", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_PAREN, Lexeme: "(", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).ToNot(BeNil())
					Expect(actual).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("Expected right parent ')' after expression"))
				})
			})
		})

		Describe("Unary", func() {
			When("its a list with a negative and number", func() {
				It("returns a tree with a unary node", func() {
					tokens := []lexer.Token{
						{Type: lexer.MINUS, Lexeme: "-", Line: 1},
						{Type: lexer.NUMBER, Literal: 5, Lexeme: "5", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Unary{
						Operator: tokens[0],
						Right:    expr.Primary{Value: 5},
					}}))
				})
			})

			When("its a list with a bang and number", func() {
				It("returns a tree with a unary node", func() {
					tokens := []lexer.Token{
						{Type: lexer.BANG, Lexeme: "!", Line: 1},
						{Type: lexer.NUMBER, Literal: 5, Lexeme: "5", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Unary{
							Operator: tokens[0],
							Right:    expr.Primary{Value: 5},
						},
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Binary{
						Left:     expr.Primary{Value: 1},
						Operator: tokens[1],
						Right:    expr.Primary{Value: 2},
					}}))
				})
			})

			When("its a list with two numbers and an a star", func() {
				It("returns a tree with one factor node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.STAR, Lexeme: "*", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Binary{
						Left:     expr.Primary{Value: 1},
						Operator: tokens[1],
						Right:    expr.Primary{Value: 2},
					}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{Expression: expr.Binary{
						Left: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 1},
						},
						Operator: tokens[3],
						Right:    expr.Primary{Value: 1},
					}}))
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list with two numbers and a minus", func() {
				It("returns a tree with one term node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.MINUS, Lexeme: "-", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
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
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left: expr.Binary{
								Left:     expr.Primary{Value: 1},
								Operator: tokens[1],
								Right:    expr.Primary{Value: 1},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: 1},
						},
					}))
				})
			})
		})

		Describe("Comparison", func() {
			When("its a list with two numbers and a greater than", func() {
				It("returns a tree with one comparison node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.GREATER, Lexeme: ">", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list with two numbers and a less than", func() {
				It("returns a tree with one comparison node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.LESS, Lexeme: "<", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list with two numbers and a greater than or equal to", func() {
				It("returns a tree with one comparison node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.GREATER_EQ, Lexeme: ">=", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list with two numbers and a less than or equal to", func() {
				It("returns a tree with one comparison node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.LESS_EQ, Lexeme: "<=", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list of multiple numbers and greater than tokens", func() {
				It("returns a nested comparison tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.GREATER, Lexeme: ">", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.GREATER, Lexeme: ">", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left: expr.Binary{
								Left:     expr.Primary{Value: 1},
								Operator: tokens[1],
								Right:    expr.Primary{Value: 1},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: 1},
						},
					}))
				})
			})
		})

		Describe("Equality", func() {
			When("its a list with two numbers and an equal", func() {
				It("returns a tree with one equality node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list with two numbers and a not equal", func() {
				It("returns a tree with one equality node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
						{Type: lexer.NUMBER, Literal: 2, Lexeme: "2", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left:     expr.Primary{Value: 1},
							Operator: tokens[1],
							Right:    expr.Primary{Value: 2},
						},
					}))
				})
			})

			When("its a list of multiple numbers and equal tokens", func() {
				It("returns a nested equality tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left: expr.Binary{
								Left:     expr.Primary{Value: 1},
								Operator: tokens[1],
								Right:    expr.Primary{Value: 1},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: 1},
						},
					}))
				})
			})

			When("its a list of multiple numbers and not equal tokens", func() {
				It("returns a nested equality tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
						{Type: lexer.NUMBER, Literal: 1, Lexeme: "1", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Binary{
							Left: expr.Binary{
								Left:     expr.Primary{Value: 1},
								Operator: tokens[1],
								Right:    expr.Primary{Value: 1},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: 1},
						},
					}))
				})
			})
		})

		Describe("Logical", func() {
			When("its a list with two true tokens and an OR", func() {
				It("returns a tree with one logical OR node", func() {
					tokens := []lexer.Token{
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.OR, Lexeme: "or", Line: 1},
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Logical{
							Left:     expr.Primary{Value: true},
							Operator: tokens[1],
							Right:    expr.Primary{Value: true},
						},
					}))
				})
			})

			When("its a list with two false tokens and an AND", func() {
				It("returns a tree with one logical AND node", func() {
					tokens := []lexer.Token{
						{Type: lexer.FALSE, Literal: "false", Lexeme: "false", Line: 1},
						{Type: lexer.AND, Lexeme: "and", Line: 1},
						{Type: lexer.FALSE, Literal: "false", Lexeme: "false", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Logical{
							Left:     expr.Primary{Value: false},
							Operator: tokens[1],
							Right:    expr.Primary{Value: false},
						},
					}))
				})
			})

			When("its a list with multiple logical OR tokens", func() {
				It("returns a nested logical OR tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.OR, Lexeme: "or", Line: 1},
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.OR, Lexeme: "or", Line: 1},
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Logical{
							Left: expr.Logical{
								Left:     expr.Primary{Value: true},
								Operator: tokens[1],
								Right:    expr.Primary{Value: true},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: true},
						},
					}))
				})
			})

			When("its a list with multiple logical AND tokens", func() {
				It("returns a nested logical AND tree node", func() {
					tokens := []lexer.Token{
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.AND, Lexeme: "and", Line: 1},
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.AND, Lexeme: "and", Line: 1},
						{Type: lexer.TRUE, Literal: "true", Lexeme: "true", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Logical{
							Left: expr.Logical{
								Left:     expr.Primary{Value: true},
								Operator: tokens[1],
								Right:    expr.Primary{Value: true},
							},
							Operator: tokens[3],
							Right:    expr.Primary{Value: true},
						},
					}))
				})
			})
		})

		Describe("Assignment", func() {
			When("its a list with a variable assignment", func() {
				It("returns an assignment expression", func() {
					tokens := []lexer.Token{
						{Type: lexer.IDENTIFIER, Lexeme: "foo", Line: 1},
						{Type: lexer.EQUAL, Lexeme: "=", Line: 1},
						{Type: lexer.NUMBER, Literal: 42, Lexeme: "42", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Expression{
						Expression: expr.Assign{
							Name:  tokens[0],
							Value: expr.Primary{Value: 42},
						},
					}))
				})
			})
		})
	})

	Describe("Statements", func() {
		Describe("Print", func() {
			When("its a list with a print token and a string", func() {
				It("returns a print statement with a string expression", func() {
					tokens := []lexer.Token{
						{Type: lexer.PRINT, Lexeme: "print", Line: 1},
						{Type: lexer.STRING, Literal: "test", Lexeme: "\"test\"", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Print{
						Expression: expr.Primary{Value: "test"},
					}))
				})
			})

			When("its a list with a print token and string but without a semicolon", func() {
				It("returns an error", func() {
					tokens := []lexer.Token{
						{Type: lexer.PRINT, Lexeme: "print", Line: 1},
						{Type: lexer.STRING, Literal: "test", Lexeme: "\"test\"", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					_, err := parser.Parse()
					Expect(err.Error()).To(ContainSubstring("reached end of file"))
				})
			})
		})

		Describe("Var", func() {
			When("its a list with a var and an identifier", func() {
				It("returns a var declaration", func() {
					tokens := []lexer.Token{
						{Type: lexer.VAR, Lexeme: "var", Line: 1},
						{Type: lexer.IDENTIFIER, Lexeme: "foo", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Var{
						Name: lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "foo", Line: 1},
					}))
				})
			})

			When("its a list with a var, identifier and initializer", func() {
				It("returns a var declaration with an initializer", func() {
					tokens := []lexer.Token{
						{Type: lexer.VAR, Lexeme: "var", Line: 1},
						{Type: lexer.IDENTIFIER, Lexeme: "foo", Line: 1},
						{Type: lexer.EQUAL, Lexeme: "=", Line: 1},
						{Type: lexer.TRUE, Lexeme: "true", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Var{
						Name:        lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "foo", Line: 1},
						Initializer: expr.Primary{Value: true},
					}))
				})
			})

			When("its a list with a var and a non-identifier token", func() {
				It("returns an error", func() {
					tokens := []lexer.Token{
						{Type: lexer.VAR, Lexeme: "var", Line: 1},
						{Type: lexer.NUMBER, Lexeme: "123", Line: 1},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).ToNot(BeNil())
					Expect(actual).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("expect variable name"))
				})
			})
		})

		Describe("Block", func() {
			When("its a list with an empty block", func() {
				It("returns a block statement with no statements", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_BRACE, Lexeme: "{", Line: 1},
						{Type: lexer.RIGHT_BRACE, Lexeme: "}", Line: 1},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 1},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Block{
						Statements: nil,
					}))
				})
			})

			When("its a list with a block containing a single statement", func() {
				It("returns a block statement with one statement", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_BRACE, Lexeme: "{", Line: 1},
						{Type: lexer.PRINT, Lexeme: "print", Line: 2},
						{Type: lexer.STRING, Literal: "test", Lexeme: "\"test\"", Line: 2},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 2},
						{Type: lexer.RIGHT_BRACE, Lexeme: "}", Line: 3},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 3},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Block{
						Statements: []stmt.Stmt{
							stmt.Print{
								Expression: expr.Primary{Value: "test"},
							},
						},
					}))
				})
			})

			When("its a list with a block containing multiple statements", func() {
				It("returns a block statement with multiple statements", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_BRACE, Lexeme: "{", Line: 1},
						{Type: lexer.PRINT, Lexeme: "print", Line: 2},
						{Type: lexer.STRING, Literal: "test1", Lexeme: "\"test1\"", Line: 2},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 2},
						{Type: lexer.PRINT, Lexeme: "print", Line: 3},
						{Type: lexer.STRING, Literal: "test2", Lexeme: "\"test2\"", Line: 3},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 3},
						{Type: lexer.RIGHT_BRACE, Lexeme: "}", Line: 4},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 4},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Block{
						Statements: []stmt.Stmt{
							stmt.Print{
								Expression: expr.Primary{Value: "test1"},
							},
							stmt.Print{
								Expression: expr.Primary{Value: "test2"},
							},
						},
					}))
				})
			})

			When("its a list with a block containing nested blocks", func() {
				It("returns a block statement with nested block statements", func() {
					tokens := []lexer.Token{
						{Type: lexer.LEFT_BRACE, Lexeme: "{", Line: 1},
						{Type: lexer.LEFT_BRACE, Lexeme: "{", Line: 2},
						{Type: lexer.PRINT, Lexeme: "print", Line: 3},
						{Type: lexer.STRING, Literal: "nested", Lexeme: "\"nested\"", Line: 3},
						{Type: lexer.SEMICOLON, Lexeme: ";", Line: 3},
						{Type: lexer.RIGHT_BRACE, Lexeme: "}", Line: 4},
						{Type: lexer.RIGHT_BRACE, Lexeme: "}", Line: 5},
						{Type: lexer.EOF, Lexeme: "EOF", Line: 5},
					}
					parser := New(tokens)
					actual, err := parser.Parse()
					Expect(err).To(BeNil())
					Expect(actual[0]).To(Equal(stmt.Block{
						Statements: []stmt.Stmt{
							stmt.Block{
								Statements: []stmt.Stmt{
									stmt.Print{
										Expression: expr.Primary{Value: "nested"},
									},
								},
							},
						},
					}))
				})
			})
		})
	})
})
