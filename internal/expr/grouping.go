package expr

type Grouping struct {
	Expression any
}

func (g Grouping) Accept(v Visitor) (any, error) {
	val, err := v.VisitGrouping(g)
	if err != nil {
		return nil, err
	}
	return val, nil
}
