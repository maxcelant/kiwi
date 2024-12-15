package stmt

type Block struct {
	Statements []Stmt
}

func (b Block) Accept(v Visitor) error {
	err := v.VisitBlockStatement(b)
	if err != nil {
		return err
	}
	return nil
}
