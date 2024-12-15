package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Assign struct {
	Name  lexer.Token
	Value Expr
}

func (a Assign) Accept(v Visitor) (any, error) {
	val, err := v.VisitAssign(a)
	if err != nil {
		return nil, err
	}
	return val, nil
}
