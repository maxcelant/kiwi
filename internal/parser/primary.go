package parser

type Primary struct {
	value any
}

func (p Primary) Accept(v Visitor) any {
	return v.VisitPrimary(p)
}
