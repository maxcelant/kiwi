package stmt

import (
	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
)

type Var struct {
	Name        lexer.Token
	Initializer expr.Expr // Can be null
}

func (vr Var) Accept(v Visitor) error {
	err := v.VisitVarDeclaration(vr)
	if err != nil {
		return err
	}
	return nil
}
