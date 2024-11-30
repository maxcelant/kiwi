package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Unary struct {
	Operator lexer.Token
	Right    any
}

func (u Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
}
