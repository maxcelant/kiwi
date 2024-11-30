package parser

import "github.com/maxcelant/kiwi/internal/lexer"

type Binary struct {
	right    any
	left     any
	operator lexer.Token
}

func (b *Binary) Accept(v Visitor) Expr {
	return v.VisitFactor(b)
}
