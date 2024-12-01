package stmt

type Visitor interface {
	VisitPrintStatement(Stmt) error
	VisitExpressionStatement(Stmt) error
}
