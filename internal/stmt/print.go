package stmt

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/expr"
)

type Print struct {
	Expression expr.Expr
}

func (p Print) Accept(v Visitor) (any, error) {
	fmt.Print(p.Expression)
	return nil, nil
}
