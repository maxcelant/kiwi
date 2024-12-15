package interpreter

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/env"
	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
	"github.com/maxcelant/kiwi/internal/stmt"
)

type Interpreter struct {
	stmts       []stmt.Stmt
	environment env.Environment
}

func New(stmts []stmt.Stmt, environment env.Environment) *Interpreter {
	return &Interpreter{
		stmts:       stmts,
		environment: environment,
	}
}

func (it *Interpreter) Interpret() {
	for _, st := range it.stmts {
		err := it.Execute(st)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (it *Interpreter) Execute(st stmt.Stmt) error {
	err := st.Accept(it)
	if err != nil {
		return err
	}
	return nil
}

func (it *Interpreter) Evaluate(ex expr.Expr) (any, error) {
	v, err := ex.Accept(it)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate expression: %w", err)
	}
	return v, nil
}

func (it *Interpreter) VisitVarDeclaration(st stmt.Stmt) error {
	var err error
	var v any
	varStmt, ok := st.(stmt.Var)
	if !ok {
		return fmt.Errorf("not an expression statement")
	}
	if varStmt.Initializer != nil {
		v, err = it.Evaluate(varStmt.Initializer)
		if err != nil {
			return err
		}
	} else {
		v = nil
	}

	it.environment.Define(varStmt.Name.Lexeme, v)
	return nil
}

func (it *Interpreter) VisitExpressionStatement(st stmt.Stmt) error {
	exprStmt, ok := st.(stmt.Expression)
	if !ok {
		return fmt.Errorf("not an expression statement")
	}
	_, err := it.Evaluate(exprStmt.Expression)
	if err != nil {
		return err
	}
	return nil
}

func (it *Interpreter) VisitPrintStatement(st stmt.Stmt) error {
	prntStmt, ok := st.(stmt.Print)
	if !ok {
		return fmt.Errorf("not an expression statement")
	}
	v, err := it.Evaluate(prntStmt.Expression)
	if err != nil {
		return err
	}
	fmt.Println(Stringify(v))
	return nil
}

func (it *Interpreter) VisitAssign(ex expr.Expr) (any, error) {
	assign, ok := ex.(expr.Assign)
	if !ok {
		return nil, fmt.Errorf("not an assign expression")
	}

	value, err := it.Evaluate(assign.Value)
	if err != nil {
		return nil, err
	}

	err = it.environment.Assign(assign.Name, value)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (it *Interpreter) VisitLogical(ex expr.Expr) (any, error) {
	logical, ok := ex.(expr.Logical)
	if !ok {
		return nil, fmt.Errorf("not a logical expression")
	}

	left, err := it.Evaluate(logical.Left)
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

	return it.Evaluate(logical.Right)
}

func (it *Interpreter) VisitBinary(ex expr.Expr) (any, error) {
	binary, ok := ex.(expr.Binary)
	if !ok {
		return nil, fmt.Errorf("not a binary expression")
	}

	left, err := it.Evaluate(binary.Left)
	if err != nil {
		return nil, err
	}

	right, err := it.Evaluate(binary.Right)
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

func (it *Interpreter) VisitUnary(ex expr.Expr) (any, error) {
	unary, ok := ex.(expr.Unary)
	if !ok {
		return nil, fmt.Errorf("not a unary expression")
	}

	right, err := it.Evaluate(unary.Right)
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

func (it *Interpreter) VisitPrimary(ex expr.Expr) (any, error) {
	primary, ok := ex.(expr.Primary)
	if !ok {
		return nil, fmt.Errorf("not a primary expression")
	}
	return primary.Value, nil
}

func (it *Interpreter) VisitVariable(ex expr.Expr) (any, error) {
	variable, ok := ex.(expr.Variable)
	if !ok {
		return nil, fmt.Errorf("not a variable expression")
	}
	val, err := it.environment.Get(variable.Name)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (it *Interpreter) VisitGrouping(ex expr.Expr) (any, error) {
	grouping, ok := ex.(expr.Grouping)
	if !ok {
		return nil, fmt.Errorf("not a grouping expression")
	}

	value, err := it.Evaluate(grouping.Expression)
	if err != nil {
		return nil, err
	}
	return value, nil
}
