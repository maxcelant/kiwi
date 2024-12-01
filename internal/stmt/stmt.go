package stmt

type Stmt interface {
	Accept(Visitor) (any, error)
}
