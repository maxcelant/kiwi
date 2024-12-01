package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Binary struct {
	Right    Expr
	Left     Expr
	Operator lexer.Token
}

func (b Binary) Accept(v Visitor) (any, error) {
	val, err := v.VisitBinary(b)
	if err != nil {
		return nil, err
	}
	return val, nil
}
