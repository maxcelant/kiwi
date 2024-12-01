package expr

type Visitor interface {
	VisitLogical(Expr) (any, error)
	VisitBinary(Expr) (any, error)
	VisitUnary(Expr) (any, error)
	VisitPrimary(Expr) (any, error)
	VisitGrouping(Expr) (any, error)
}
