package stmt

import "github.com/maxcelant/kiwi/internal/expr"

type Expression struct {
	Expression expr.Expr
}

func (e Expression) Accept(v Visitor) error {
	err := v.VisitExpressionStatement(e)
	if err != nil {
		return err
	}
	return nil
}
