package interpreter

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/parser"
)

type Interpreter struct {
	expr parser.Expr
}

func New(expr parser.Expr) *Interpreter {
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

func (it *Interpreter) VisitBinary(expr expr.Expr) any {

}

func (it *Interpreter) VisitUnary(expr expr.Expr) any {

}

func (it *Interpreter) VisitPrimary(expr expr.Expr) any {
	primary, ok := expr.(*expr.Primary)
	if !ok {
		// handle error
	}
	return primary.Value
}

func (it *Interpreter) VisitGrouping(expr expr.Expr) any {

}

func (it *Interpreter) Stringify(obj any) (string, error) {

}
