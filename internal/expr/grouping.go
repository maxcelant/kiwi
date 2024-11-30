package expr

type Grouping struct {
	Expr any
}

func (g Grouping) Accept(v Visitor) any {
	return v.VisitGrouping(g)
}
