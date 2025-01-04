package stmt

import "github.com/maxcelant/kiwi/internal/expr"

type If struct {
  Condition expr.Expr 
  ThenBranch Stmt
  ElseBranch Stmt
}

func (i *If) Accept(v Visitor) error {
	err := v.VisitIfStatement(i)
	if err != nil {
		return err
	}
	return nil
}
