package expr

import "github.com/maxcelant/kiwi/internal/lexer"

type Logical struct {
	Right    any
	Left     any
	Operator lexer.Token
}

func (l Logical) Accept(v Visitor) (any, error) {
	val, err := v.VisitLogical(l)
	if err != nil {
		return nil, err
	}
	return val, nil
}
