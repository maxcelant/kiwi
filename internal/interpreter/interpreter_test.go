package interpreter

import (
	"testing"

	"github.com/maxcelant/kiwi/internal/env"
	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
	"github.com/maxcelant/kiwi/internal/stmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInterpreter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interpreter Suite")
}

var it *Interpreter

var _ = Describe("Interpreter", func() {
	BeforeEach(func() {
		it = New(nil, env.New())
	})

	Describe("Visit Primary Expr", func() {
		When("the parse tree has a single primary number node", func() {
			It("should return the value", func() {
				node := expr.Primary{Value: 1}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(1))
			})
		})

		When("the parse tree has a single primary nil node", func() {
			It("should return nil", func() {
				node := expr.Primary{Value: nil}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(BeNil())
			})
		})

		When("the parse tree has a single primary string node", func() {
			It("should return the string value", func() {
				node := expr.Primary{Value: "test"}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal("test"))
			})
		})
	})

	Describe("Visit Grouping Expr", func() {
		When("the parse tree has a grouping node with a primary number node", func() {
			It("should return the value of the primary node", func() {
				primaryNode := expr.Primary{Value: 1}
				groupingNode := expr.Grouping{Expression: primaryNode}
				actual, err := it.Evaluate(groupingNode)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(1))
			})
		})

		When("the parse tree has a grouping node with a primary string node", func() {
			It("should return the value of the primary node", func() {
				primaryNode := expr.Primary{Value: "test"}
				groupingNode := expr.Grouping{Expression: primaryNode}
				actual, err := it.Evaluate(groupingNode)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal("test"))
			})
		})
	})

	Describe("Visit Unary Expr", func() {
		When("the parse tree has a unary node that includes a bang and a true", func() {
			It("should return the opposite of that value", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.BANG, Lexeme: "!", Line: 1},
					Right:    expr.Primary{Value: true},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a unary node that includes a bang and a false", func() {
			It("should return the opposite of that value", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.BANG, Lexeme: "!", Line: 1},
					Right:    expr.Primary{Value: false},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})
		})

		When("the parse tree has a unary node that includes a bang and a nil", func() {
			It("should return true", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.BANG, Lexeme: "!", Line: 1},
					Right:    expr.Primary{Value: nil},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})
		})

		When("the parse tree has a unary node that includes a bang and a non-nil value", func() {
			It("should return false", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.BANG, Lexeme: "!", Line: 1},
					Right:    expr.Primary{Value: "non-nil"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a unary node that includes a minus and a positive number", func() {
			It("should return the negative of that number", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(-5))
			})
		})

		When("the parse tree has a unary node that includes a minus and a negative number", func() {
			It("should return the positive of that number", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: -5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(5))
			})
		})

		When("the parse tree has a unary node that includes a minus and zero", func() {
			It("should return zero", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: 0},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(0))
			})
		})

		When("the parse tree has a unary node that includes a minus and a non-number value", func() {
			It("should return an error", func() {
				node := expr.Unary{
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operand must be a number"))
			})
		})
	})

	Describe("Visit Binary Expr", func() {
		When("the parse tree has a binary node that adds two numbers", func() {
			It("should return the sum of those numbers", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(8))
			})
		})

		When("the parse tree has a binary node that adds two strings", func() {
			It("should return the concatenation of those strings", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: "world"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal("helloworld"))
			})
		})

		When("the parse tree has a binary node with a non-string left operand for addition", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: " world"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must both be a numbers or strings for add operation"))
			})
		})

		When("the parse tree has a binary node with a non-string right operand for addition", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must both be a numbers or strings for add operation"))
			})
		})

		When("the parse tree has a binary node with a non-number left operand", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "non-number"},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must both be a numbers or strings for add operation"))
			})
		})

		When("the parse tree has a binary node with a non-number right operand", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must both be a numbers or strings for add operation"))
			})
		})

		When("the parse tree has a binary node that subtracts two numbers", func() {
			It("should return the difference of those numbers", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: 4},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(6))
			})
		})

		When("the parse tree has a binary node that subtracts a smaller number from a larger number", func() {
			It("should return the negative difference of those numbers", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: 10},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(-7))
			})
		})

		When("the parse tree has a binary node with a non-number left operand for subtraction", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "non-number"},
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: 4},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node with a non-number right operand for subtraction", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node that multiplies two numbers", func() {
			It("should return the product of those numbers", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.STAR, Lexeme: "*", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(15))
			})
		})

		When("the parse tree has a binary node with a non-number left operand for multiplication", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "non-number"},
					Operator: lexer.Token{Type: lexer.STAR, Lexeme: "*", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node with a non-number right operand for multiplication", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.STAR, Lexeme: "*", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node that divides two numbers", func() {
			It("should return the quotient of those numbers", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.SLASH, Lexeme: "/", Line: 1},
					Right:    expr.Primary{Value: 2},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(5))
			})
		})

		When("the parse tree has a binary node with a non-number left operand for division", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "non-number"},
					Operator: lexer.Token{Type: lexer.SLASH, Lexeme: "/", Line: 1},
					Right:    expr.Primary{Value: 2},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node with a non-number right operand for division", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.SLASH, Lexeme: "/", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node that divides a number by zero", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.SLASH, Lexeme: "/", Line: 1},
					Right:    expr.Primary{Value: 0},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("division by zero"))
			})
		})

		When("the parse tree has a binary node that compares two numbers with greater than", func() {
			It("should return true if the left operand is greater than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.GREATER, Lexeme: ">", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the left operand is not greater than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.GREATER, Lexeme: ">", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node with a non-number left operand for greater than", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "non-number"},
					Operator: lexer.Token{Type: lexer.GREATER, Lexeme: ">", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node with a non-number right operand for greater than", func() {
			It("should return an error", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.GREATER, Lexeme: ">", Line: 1},
					Right:    expr.Primary{Value: "non-number"},
				}
				_, err := it.Evaluate(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must be a number"))
			})
		})

		When("the parse tree has a binary node that compares two numbers with greater than or equal", func() {
			It("should return true if the left operand is greater than or equal to the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.GREATER_EQ, Lexeme: ">=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return true if the left operand is equal to the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.GREATER_EQ, Lexeme: ">=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the left operand is less than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.GREATER_EQ, Lexeme: ">=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two numbers with less than", func() {
			It("should return true if the left operand is less than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.LESS, Lexeme: "<", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the left operand is not less than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.LESS, Lexeme: "<", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two numbers with less than or equal", func() {
			It("should return true if the left operand is less than or equal to the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 3},
					Operator: lexer.Token{Type: lexer.LESS_EQ, Lexeme: "<=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return true if the left operand is equal to the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.LESS_EQ, Lexeme: "<=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the left operand is greater than the right operand", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 10},
					Operator: lexer.Token{Type: lexer.LESS_EQ, Lexeme: "<=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two numbers for equality", func() {
			It("should return true if the numbers are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the numbers are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: 3},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two booleans for equality", func() {
			It("should return true if the booleans are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: true},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the booleans are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: false},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two booleans for inequality", func() {
			It("should return true if the booleans are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: false},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the booleans are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: true},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two numbers for inequality", func() {
			It("should return true if the numbers are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: 3},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the numbers are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: 5},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: 5},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two strings for equality", func() {
			It("should return true if the strings are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: "hello"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the strings are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.EQUAL_EQUAL, Lexeme: "==", Line: 1},
					Right:    expr.Primary{Value: "world"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a binary node that compares two strings for inequality", func() {
			It("should return true if the strings are not equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: "world"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})

			It("should return false if the strings are equal", func() {
				node := expr.Binary{
					Left:     expr.Primary{Value: "hello"},
					Operator: lexer.Token{Type: lexer.BANG_EQ, Lexeme: "!=", Line: 1},
					Right:    expr.Primary{Value: "hello"},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})
	})

	Describe("Visit Logical Expr", func() {
		When("the parse tree has a logical AND node with both operands true", func() {
			It("should return true", func() {
				node := expr.Logical{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.AND, Lexeme: "and", Line: 1},
					Right:    expr.Primary{Value: true},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})
		})

		When("the parse tree has a logical AND node with one operand false", func() {
			It("should return false", func() {
				node := expr.Logical{
					Left:     expr.Primary{Value: true},
					Operator: lexer.Token{Type: lexer.AND, Lexeme: "and", Line: 1},
					Right:    expr.Primary{Value: false},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a logical OR node with both operands false", func() {
			It("should return false", func() {
				node := expr.Logical{
					Left:     expr.Primary{Value: false},
					Operator: lexer.Token{Type: lexer.OR, Lexeme: "or", Line: 1},
					Right:    expr.Primary{Value: false},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(false))
			})
		})

		When("the parse tree has a logical OR node with one operand true", func() {
			It("should return true", func() {
				node := expr.Logical{
					Left:     expr.Primary{Value: false},
					Operator: lexer.Token{Type: lexer.OR, Lexeme: "or", Line: 1},
					Right:    expr.Primary{Value: true},
				}
				actual, err := it.Evaluate(node)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(true))
			})
		})
	})

	Describe("Visit Expression Stmt", func() {
		When("the parse tree has an expression statement with a primary number node", func() {
			It("should evaluate the expression without error", func() {
				node := stmt.Expression{Expression: expr.Primary{Value: 1}}
				err := it.Execute(node)
				Expect(err).To(BeNil())
			})
		})

		When("the parse tree has an expression statement with a primary string node", func() {
			It("should evaluate the expression without error", func() {
				node := stmt.Expression{Expression: expr.Primary{Value: "test"}}
				err := it.Execute(node)
				Expect(err).To(BeNil())
			})
		})

		When("the parse tree has an expression statement with a binary addition node", func() {
			It("should evaluate the expression without error", func() {
				node := stmt.Expression{
					Expression: expr.Binary{
						Left:     expr.Primary{Value: 3},
						Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
						Right:    expr.Primary{Value: 5},
					},
				}
				err := it.Execute(node)
				Expect(err).To(BeNil())
			})
		})

		When("the parse tree has an expression statement with a unary minus node", func() {
			It("should evaluate the expression without error", func() {
				node := stmt.Expression{
					Expression: expr.Unary{
						Operator: lexer.Token{Type: lexer.MINUS, Lexeme: "-", Line: 1},
						Right:    expr.Primary{Value: 5},
					},
				}
				err := it.Execute(node)
				Expect(err).To(BeNil())
			})
		})

		When("the parse tree has an expression statement with a logical AND node", func() {
			It("should evaluate the expression without error", func() {
				node := stmt.Expression{
					Expression: expr.Logical{
						Left:     expr.Primary{Value: true},
						Operator: lexer.Token{Type: lexer.AND, Lexeme: "and", Line: 1},
						Right:    expr.Primary{Value: true},
					},
				}
				err := it.Execute(node)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Visit Print Stmt", func() {

	})

	Describe("Var Declaration", func() {
		When("the parse tree has a var declaration with an initializer", func() {
			It("should define the variable in the environment", func() {
				node := stmt.Var{
					Name:        lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "x", Line: 1},
					Initializer: expr.Primary{Value: 42},
				}
				err := it.VisitVarDeclaration(node)
				Expect(err).To(BeNil())
				value, err := it.environment.Get(lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "x"})
				Expect(err).To(BeNil())
				Expect(value).To(Equal(42))
			})
		})

		When("the parse tree has a var declaration without an initializer", func() {
			It("should define the variable with a nil value in the environment", func() {
				node := stmt.Var{
					Name:        lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "y", Line: 1},
					Initializer: nil,
				}
				err := it.VisitVarDeclaration(node)
				Expect(err).To(BeNil())
				value, err := it.environment.Get(lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "y"})
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})

		When("the parse tree has a var declaration with an invalid initializer", func() {
			It("should return an error", func() {
				node := stmt.Var{
					Name: lexer.Token{Type: lexer.IDENTIFIER, Lexeme: "z", Line: 1},
					Initializer: expr.Binary{
						Left:     expr.Primary{Value: 1},
						Operator: lexer.Token{Type: lexer.PLUS, Lexeme: "+", Line: 1},
						Right:    expr.Primary{Value: "invalid"},
					},
				}
				err := it.VisitVarDeclaration(node)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("operands must both be a numbers or strings for add operation"))
			})
		})
	})
})
