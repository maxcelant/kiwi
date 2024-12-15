package expr

type Visitor interface {
	VisitAssign(Expr) (any, error)
	VisitVariable(Expr) (any, error)
	VisitLogical(Expr) (any, error)
	VisitBinary(Expr) (any, error)
	VisitUnary(Expr) (any, error)
	VisitPrimary(Expr) (any, error)
	VisitGrouping(Expr) (any, error)
}
