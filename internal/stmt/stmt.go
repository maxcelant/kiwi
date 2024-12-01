package stmt

type Stmt interface {
	Accept(Visitor) error
}
