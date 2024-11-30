package parser

type Grouping struct {
	expr any
}

func (g Grouping) Accept(v Visitor) Expr {
	return v.VisitGrouping(g)
}
