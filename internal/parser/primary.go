package parser

type Primary struct {
	Value any
}

func (p Primary) Accept(v Visitor) any {
	return v.VisitPrimary(p)
}
