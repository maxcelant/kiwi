package stmt

type Visitor interface {
	VisitVarDeclaration(Stmt) error
	VisitPrintStatement(Stmt) error
	VisitExpressionStatement(Stmt) error
}
