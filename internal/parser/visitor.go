package parser

type Visitor interface {
	VisitFactor(Expr) Expr
	VisitUnary(Expr) Expr
	VisitPrimary(Expr) Expr
	VisitGrouping(Expr) Expr
}
