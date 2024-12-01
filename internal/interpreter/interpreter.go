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
	str, err = Stringify(obj)
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

func (it *Interpreter) VisitLogical(expr exp.Expr) (any, error) {
	logical, ok := expr.(exp.Logical)
	if !ok {
		return nil, fmt.Errorf("not a logical expression")
	}

	rightExpr, ok := logical.Right.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("unary.Right is not a type exp.Expr")
	}

	leftExpr, ok := logical.Left.(exp.Expr)
	if !ok {
		return nil, fmt.Errorf("unary.Left is not a type exp.Expr")
	}

	left, err := it.Evaluate(leftExpr)
	if err != nil {
		return nil, err
	}

	// We only evaluate the right operand if the left one is falsey
	// If the left is truthy, we just return that one
	if logical.Operator.Type == lexer.OR {
		if IsTruthy(left) {
			return left, nil
		}
	} else {
		if !IsTruthy(left) {
			return left, nil
		}
	}
	return it.Evaluate(rightExpr)

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

	if binary.Operator.Type == lexer.EQUAL_EQUAL {
		if ok := Compare(left, right, WithInt(), WithBool(), WithString()); !ok {
			return nil, fmt.Errorf("operands must be a number or boolean for equality operation")
		}
		return left == right, nil
	}

	if binary.Operator.Type == lexer.BANG_EQ {
		if ok := Compare(left, right, WithInt(), WithBool(), WithString()); !ok {
			return nil, fmt.Errorf("operands must be a number or boolean for inequality operation")
		}
		return left != right, nil
	}

	if binary.Operator.Type == lexer.GREATER {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for greater than operation")
		}
		return left.(int) > right.(int), nil
	}

	if binary.Operator.Type == lexer.GREATER_EQ {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for greater than or equal operation")
		}
		return left.(int) >= right.(int), nil
	}

	if binary.Operator.Type == lexer.LESS {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for less than operation")
		}
		return left.(int) < right.(int), nil
	}

	if binary.Operator.Type == lexer.LESS_EQ {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for less than or equal operation")
		}
		return left.(int) <= right.(int), nil
	}

	if binary.Operator.Type == lexer.PLUS {
		if ok := Compare(left, right, WithInt(), WithString()); !ok {
			return nil, fmt.Errorf("operands must both be a numbers or strings for add operation")
		}

		switch left := left.(type) {
		case int:
			return left + right.(int), nil
		case string:
			return left + right.(string), nil
		default:
			return nil, fmt.Errorf("operands must be either both numbers or both strings for add operation")
		}
	}

	if binary.Operator.Type == lexer.MINUS {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for subtract operation")
		}
		return left.(int) - right.(int), nil
	}

	if binary.Operator.Type == lexer.SLASH {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for division operation")
		}
		if right.(int) == 0 {
			return nil, fmt.Errorf("cannot perform division by zero")
		}
		return left.(int) / right.(int), nil
	}

	if binary.Operator.Type == lexer.STAR {
		if ok := Compare(left, right, WithInt()); !ok {
			return nil, fmt.Errorf("operands must be a number for multiplication operation")
		}
		return left.(int) * right.(int), nil
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
		return !IsTruthy(right), nil
	}

	if unary.Operator.Type == lexer.MINUS {
		num, ok := isNumber(right)
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
