package parser

type Primary struct {
	value interface{}
}

func (p *Primary) Accept(v Visitor) Expr {
	return v.VisitPrimary(p)
}
