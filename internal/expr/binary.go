package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Binary struct {
	Right    any
	Left     any
	Operator lexer.Token
}

func (b Binary) Accept(v Visitor) any {
	return v.VisitBinary(b)
}
