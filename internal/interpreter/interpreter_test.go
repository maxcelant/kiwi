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

var itr *Interpreter

var _ = Describe("Interpreter", func() {
	Describe("Visit Primary", func() {
		When("the parse tree has a single primary number node", func() {
			It("should return the value", func() {
				node := expr.Primary{Value: 1}
				itr = New(node)
				actual := itr.Evaluate()
				Expect(actual).To(Equal(1))
			})
		})

		When("the parse tree has a single primary nil node", func() {
			It("should return nil", func() {
				node := expr.Primary{Value: nil}
				itr = New(node)
				actual := itr.Evaluate()
				Expect(actual).To(BeNil())
			})
		})

		When("the parse tree has a single primary string node", func() {
			It("should return the string value", func() {
				node := expr.Primary{Value: "test"}
				itr = New(node)
				actual := itr.Evaluate()
				Expect(actual).To(Equal("test"))
			})
		})
	})
})
