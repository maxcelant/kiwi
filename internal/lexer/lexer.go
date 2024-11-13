package lexer

type Lexer struct {
	// globalPos int64
	start    int64
	curr     int64
	currLine int64
	tokens   []Token
}

func String(ch byte) string {
	return string(ch)
}

func (l *Lexer) ScanLine(line string) ([]Token, error) {
	l.currLine += 1
	l.curr = 0
	l.tokens = []Token{}
	if line == "" {
		return l.tokens, nil
	}

	for {
		l.start = l.curr
		if l.curr >= int64(len(line)) {
			break
		}
		l.scanToken(line)
	}

	return l.tokens, nil
}

func (l *Lexer) scanToken(line string) {
	ch := line[l.curr]
	var token Token

	switch ch {
	case ';':
		token = l.addToken(line, SEMICOLON)
	}

	l.tokens = append(l.tokens, token)
	l.advance()
}

func (l *Lexer) addToken(line string, tokenType TokenType) Token {
	ch := line[l.start : l.curr+1]
	token := Token{
		Type:    tokenType,
		Literal: ch,
		Lexeme:  ch,
		Line:    l.currLine,
	}
	return token
}

// func (l *Lexer) addToken(line, tokenType TokenType, literal interface{}) {
// 	token := Token{
// 		Type:    tokenType,
// 		Literal: literal,
// 		Lexeme:  String(ch),
// 		Line:    l.currLine,
// 	}

// }

func (l *Lexer) advance() {
	l.curr += 1
}
