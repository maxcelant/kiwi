package stmt

type Visitor interface {
	VisitPrintStatement(Stmt) (any, error)
	VisitExpressionStatement(Stmt) (any, error)
}
