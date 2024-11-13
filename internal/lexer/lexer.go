package lexer

type Lexer struct {
	// globalPos int64
	start    int64
	curr     int64
	currLine int64
	tokens   []Token
	line     string
}

func String(ch byte) string {
	return string(ch)
}

func (l *Lexer) ScanLine(line string) ([]Token, error) {
	l.line = line
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
		l.scanToken()
	}

	return l.tokens, nil
}

func (l *Lexer) scanToken() {
	ch := l.line[l.curr]
	var token Token

	switch ch {
	case ';':
		token = l.addToken(SEMICOLON)
	case '=':
		next := l.match('=')
		if next {
			l.advance()
			token = l.addToken(EQUAL_EQUAL)
		} else {
			token = l.addToken(EQUAL)
		}
	}

	l.tokens = append(l.tokens, token)
	l.advance()
}

func (l *Lexer) addToken(tokenType TokenType) Token {
	ch := l.line[l.start : l.curr+1]
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

func (l *Lexer) match(next byte) (matches bool) {
	ch := l.peek()
	return ch == next
}

func (l *Lexer) peek() (next byte) {
	if l.curr+1 >= int64(len(l.line)) {
		return 0
	}
	return l.line[l.curr+1]
}
