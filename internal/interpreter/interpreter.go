package interpreter

import (
	"fmt"

	exp "github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
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
	obj, err := it.Evaluate(it.expr)
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

func (it *Interpreter) Evaluate(expr exp.Expr) (any, error) {
	v, err := expr.Accept(it)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate expression: %w", err)
	}
	return v, nil
}

func (it *Interpreter) VisitBinary(expr exp.Expr) (any, error) {
	return "", nil
}

func (it *Interpreter) VisitUnary(expr exp.Expr) (any, error) {
	unary, ok := expr.(exp.Unary)
	if !ok {
		return nil, fmt.Errorf("not a unary expression")
	}

	rightExpr, ok := unary.Right.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("unary.Right is not of type exp.Expr")
	}

	right, err := it.Evaluate(rightExpr)
	if err != nil {
		return nil, err
	}

	if unary.Operator.Type == lexer.BANG {
		return !it.IsTruthy(right), nil
	}

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
	value, err := it.Evaluate(expression)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (it *Interpreter) Stringify(obj any) (string, error) {
	return "", nil
}

func (it *Interpreter) IsTruthy(v any) bool {
	return true
}
