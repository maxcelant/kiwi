package parser

type Visitor interface {
	VisitPrimary(Expr) Expr
}
