package interpreter

import (
	"testing"

	"github.com/maxcelant/kiwi/internal/expr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInterpreter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interpreter Suite")
}

var it *Interpreter

var _ = Describe("Interpreter", func() {
	Describe("Visit Primary", func() {
		When("the parse tree has a single primary number node", func() {
			It("should return the value", func() {
				node := expr.Primary{Value: 1}
				it = New(node)
				actual, err := it.Evaluate()
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(1))
			})
		})

		When("the parse tree has a single primary nil node", func() {
			It("should return nil", func() {
				node := expr.Primary{Value: nil}
				it = New(node)
				actual, err := it.Evaluate()
				Expect(err).To(BeNil())
				Expect(actual).To(BeNil())
			})
		})

		When("the parse tree has a single primary string node", func() {
			It("should return the string value", func() {
				node := expr.Primary{Value: "test"}
				it = New(node)
				actual, err := it.Evaluate()
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
				it = New(groupingNode)
				actual, err := it.Evaluate()
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(1))
			})
		})

		When("the parse tree has a grouping node with a primary string node", func() {
			It("should return the value of the primary node", func() {
				primaryNode := expr.Primary{Value: "test"}
				groupingNode := expr.Grouping{Expression: primaryNode}
				it = New(groupingNode)
				actual, err := it.Evaluate()
				Expect(err).To(BeNil())
				Expect(actual).To(Equal("test"))
			})
		})
	})
})
