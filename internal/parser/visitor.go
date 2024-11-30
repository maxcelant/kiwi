package parser

type Visitor interface {
	VisitBinary(Expr) Expr
	VisitUnary(Expr) Expr
	VisitPrimary(Expr) Expr
	VisitGrouping(Expr) Expr
}
