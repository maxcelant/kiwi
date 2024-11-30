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
	binary, ok := expr.(exp.Binary)
	if !ok {
		return nil, fmt.Errorf("not a binary expression")
	}

	rightExpr, ok := binary.Right.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("unary.Right is not a type exp.Expr")
	}

	leftExpr, ok := binary.Left.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("unary.Left is not a type exp.Expr")
	}

	left, err := it.Evaluate(leftExpr)
	if err != nil {
		return nil, err
	}

	right, err := it.Evaluate(rightExpr)
	if err != nil {
		return nil, err
	}

	if binary.Operator.Type == lexer.PLUS {
		l, r, ok := it.BothOperandsNumbers(left, right)
		if !ok {
			return nil, fmt.Errorf("operands must be a number")
		}
		return l + r, nil
	}

	if binary.Operator.Type == lexer.MINUS {
		l, r, ok := it.BothOperandsNumbers(left, right)
		if !ok {
			return nil, fmt.Errorf("operands must be a number")
		}
		return l - r, nil
	}

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

	if unary.Operator.Type == lexer.MINUS {
		num, ok := it.isNumber(right)
		if !ok {
			return nil, fmt.Errorf("operand must be a number")
		}
		return -num, nil
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
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}

func (it *Interpreter) isNumber(v any) (int, bool) {
	switch v.(type) {
	case int, float64:
		return v.(int), true
	default:
		return 0.0, false
	}
}

func (it *Interpreter) BothOperandsNumbers(a any, b any) (int, int, bool) {
	left, ok := it.isNumber(a)
	if !ok {
		return 0, 0, false
	}
	right, ok := it.isNumber(b)
	if !ok {
		return 0, 0, false
	}
	return left, right, true
}

func (it *Interpreter) isString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}
