package stmt

type Visitor interface {
  VisitIfStatement(Stmt) error
	VisitBlockStatement(Stmt) error
	VisitVarDeclaration(Stmt) error
	VisitPrintStatement(Stmt) error
	VisitExpressionStatement(Stmt) error
}
