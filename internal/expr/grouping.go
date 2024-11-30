package expr

type Grouping struct {
	Expr any
}

func (g Grouping) Accept(v Visitor) (any, error) {
	val, err := v.VisitGrouping(g)
	if err != nil {
		return nil, err
	}
	return val, nil
}
