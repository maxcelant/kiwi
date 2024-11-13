package lexer

type Lexer struct {
	// globalPos int64
	start    int64
	curr     int64
	currLine int64
	tokens   []Token
	line     string
}

func isAlpha(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// func isNumber(b byte) bool {
// 	return (b >= '0' && b <= '9')
// }

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

	if ch == ' ' || ch == '\r' || ch == '\t' {
		l.advance()
		return
	} else if ch == ';' {
		token = l.addToken(SEMICOLON)
	} else if ch == '=' {
		next := l.match('=')
		if next {
			l.advance()
			token = l.addToken(EQUAL_EQUAL)
		} else {
			token = l.addToken(EQUAL)
		}
	} else if ch == '"' {
		token = l.handleString()
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

func (l *Lexer) addTokenWithLiteral(tokenType TokenType, literal interface{}) Token {
	ch := l.line[l.start : l.curr+1]
	token := Token{
		Type:    tokenType,
		Literal: literal,
		Lexeme:  ch,
		Line:    l.currLine,
	}
	return token
}

func (l *Lexer) handleString() Token {
	for {
		next := l.peek()
		if next == 0 || next == '"' {
			break
		}
		if isAlpha(next) {
			l.advance()
		}
	}
	l.advance()
	return l.addTokenWithLiteral(STRING, l.line[l.start+1:l.curr])
}

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
