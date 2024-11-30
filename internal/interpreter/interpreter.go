package interpreter

import (
	"fmt"

	exp "github.com/maxcelant/kiwi/internal/expr"
)

type Interpreter struct {
	expr exp.Expr
}

func New(expr exp.Expr) *Interpreter {
	return &Interpreter{
		expr: expr,
	}
}

func (it *Interpreter) Interpret() {
	var str string
	obj, err := it.Evaluate()
	if err != nil {
		fmt.Println(err)
		return
	}
	str, err = it.Stringify(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

func (it *Interpreter) Evaluate() (any, error) {
	v, err := it.expr.Accept(it)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate expression: %w", err)
	}
	return v, nil
}

func (it *Interpreter) VisitBinary(expr exp.Expr) (any, error) {
	return "", nil
}

func (it *Interpreter) VisitUnary(expr exp.Expr) (any, error) {
	return "", nil
}

func (it *Interpreter) VisitPrimary(expr exp.Expr) (any, error) {
	primary, ok := expr.(exp.Primary)
	if !ok {
		return nil, fmt.Errorf("not a primary expression")
	}
	return primary.Value, nil
}

func (it *Interpreter) VisitGrouping(expr exp.Expr) (any, error) {
	grouping, ok := expr.(exp.Grouping)
	if !ok {
		return nil, fmt.Errorf("not a grouping expression")
	}
	expression, ok := grouping.Expression.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("not a grouping expression")
	}
	return expression.Accept(it)
}

func (it *Interpreter) Stringify(obj any) (string, error) {
	return "", nil
}
