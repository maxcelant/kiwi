package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Variable struct {
	Name lexer.Token
}

func (vr Variable) Accept(v Visitor) (any, error) {
	val, err := v.VisitVariable(vr)
	if err != nil {
		return nil, err
	}
	return val, nil
}
