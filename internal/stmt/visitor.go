package stmt

// Statements perform "side-effects" and don't actually evaluate
// to a value, hence why we don't return a value for this interface
type Visitor interface {
	VisitPrintStatement(Stmt) error
	VisitExpressionStatement(Stmt) error
}
