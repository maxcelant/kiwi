package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Binary struct {
	Right    any
	Left     any
	Operator lexer.Token
}

func (b Binary) Accept(v Visitor) (any, error) {
	val, err := v.VisitBinary(b)
	if err != nil {
		return nil, err
	}
	return val, nil
}
