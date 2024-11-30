package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Unary struct {
	Operator lexer.Token
	Right    any
}

func (u Unary) Accept(v Visitor) (any, error) {
	val, err := v.VisitUnary(u)
	if err != nil {
		return nil, err
	}
	return val, nil
}
