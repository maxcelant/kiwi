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
	obj := it.Evaluate()
	str, err := it.Stringify(obj)
	if err != nil {
		// handle error
	}
	fmt.Println(str)
}

func (it *Interpreter) Evaluate() any {
	return it.expr.Accept(it)
}

func (it *Interpreter) VisitBinary(expr exp.Expr) any {
	return ""
}

func (it *Interpreter) VisitUnary(expr exp.Expr) any {
	return ""
}

func (it *Interpreter) VisitPrimary(expr exp.Expr) any {
	primary, ok := expr.(exp.Primary)
	if !ok {
		// handle error
	}
	return primary.Value
}

func (it *Interpreter) VisitGrouping(expr exp.Expr) any {
	return ""
}

func (it *Interpreter) Stringify(obj any) (string, error) {
	return "", nil
}
