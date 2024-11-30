package expr

type Primary struct {
	Value any
}

func (p Primary) Accept(v Visitor) (any, error) {
	val, err := v.VisitPrimary(p)
	if err != nil {
		return nil, err
	}
	return val, nil
}
