package lexer

type Lexer struct {
	GlobalPosition  int64
	CurrentPosition int64
}

func (l *Lexer) ScanLine(line string) ([]Token, error) {
	l.CurrentPosition = 0
	if line == "" {
		return []Token{}, nil
	}

	return []Token{}, nil
}
