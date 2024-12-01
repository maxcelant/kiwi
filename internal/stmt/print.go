package stmt

import (
	"github.com/maxcelant/kiwi/internal/expr"
)

type Print struct {
	Expression expr.Expr
}

func (p Print) Accept(v Visitor) error {
	err := v.VisitPrintStatement(p)
	if err != nil {
		return err
	}
	return nil
}
