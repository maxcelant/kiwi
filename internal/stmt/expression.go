package stmt

import "github.com/maxcelant/kiwi/internal/expr"

type Expression struct {
	Expression expr.Expr
}

func (e Expression) Accept(v Visitor) (any, error) {
	v.VisitExpressionStatement(e)
	return nil, nil
}
