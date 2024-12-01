package stmt

import (
	"github.com/maxcelant/kiwi/internal/expr"
)

type Print struct {
	Expression expr.Expr
}

func (p Print) Accept(v Visitor) (any, error) {
	v.VisitPrintStatement(p)
	return nil, nil
}
