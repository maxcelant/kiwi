package expr

type Visitor interface {
	VisitBinary(Expr) any
	VisitUnary(Expr) any
	VisitPrimary(Expr) any
	VisitGrouping(Expr) any
}
