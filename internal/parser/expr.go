package parser

type Expr interface {
	Accept(Visitor) Expr
}
