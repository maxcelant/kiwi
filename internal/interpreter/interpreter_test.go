package interpreter

import (
	"testing"

	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
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
		it = New(nil)
	})

	Describe("Visit Primary", func() {
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

	Describe("Visit Grouping", func() {
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

	Describe("Visit Unary", func() {
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
	})
})
