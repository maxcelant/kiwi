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

var lexer Lexer

var _ = Describe("Lexer", func() {

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
})
