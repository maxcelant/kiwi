package parser

type Grouping struct {
	expr any
}

func (g Grouping) Accept(v Visitor) any {
	return v.VisitGrouping(g)
}
