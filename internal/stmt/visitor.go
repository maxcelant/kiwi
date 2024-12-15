package stmt

type Visitor interface {
	VisitBlockStatement(Stmt) error
	VisitVarDeclaration(Stmt) error
	VisitPrintStatement(Stmt) error
	VisitExpressionStatement(Stmt) error
}
