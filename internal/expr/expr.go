package expr

type Expr interface {
	Accept(Visitor) (any, error)
}
