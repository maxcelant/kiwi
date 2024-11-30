package interpreter

import (
	"fmt"

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
	str, err := it.Stringify()
	if err != nil {
		// handle error
	}
	fmt.Println(str)
}

func (it *Interpreter) Evaluate() any {
	return it.expr.Accept(it)
}

func (it *Interpreter) VisitBinary(expr parser.Expr) parser.Expr {

}

func (it *Interpreter) VisitUnary(expr parser.Expr) parser.Expr {

}

func (it *Interpreter) VisitPrimary(expr parser.Expr) parser.Expr {

}

func (it *Interpreter) VisitGrouping(expr parser.Expr) parser.Expr {

}

func (it *Interpreter) Stringify() (string, error) {

}
