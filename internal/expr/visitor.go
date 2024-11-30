package expr

type Visitor interface {
	VisitBinary(Expr) (any, error)
	VisitUnary(Expr) (any, error)
	VisitPrimary(Expr) (any, error)
	VisitGrouping(Expr) (any, error)
}
