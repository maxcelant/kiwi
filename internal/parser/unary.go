package parser

import "github.com/maxcelant/kiwi/internal/lexer"

type Unary struct {
	operator lexer.Token
	right    any
}

func (u Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
}
