package parser

type Primary struct {
	value any
}

func (p Primary) Accept(v Visitor) Expr {
	return v.VisitPrimary(p)
}
