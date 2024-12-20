package lexer

import (
	"errors"
)

type Lexer struct {
	Line   int64
	start  int64
	curr   int64
	tokens []Token
	source string
}

var keywords = map[string]TokenType{
	"if":     IF,
	"else":   ELSE,
	"or":     OR,
	"and":    AND,
	"for":    FOR,
	"while":  WHILE,
	"return": RETURN,
	"fn":     FUNC,
	"class":  CLASS,
	"var":    VAR,
	"print":  PRINT,
	"true":   TRUE,
	"false":  FALSE,
	"nil":    NIL,
}

func New(source string) *Lexer {
	return &Lexer{
		source: source,
		Line:   0,
		start:  0,
		curr:   0,
		tokens: []Token{},
	}
}

func (l *Lexer) Scan() ([]Token, error) {
	for !l.atEnd() {
		l.start = l.curr
		err := l.scanToken()
		if err != nil {
			return nil, err
		}
	}
	l.tokens = append(l.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: l.Line})
	return l.tokens, nil
}

func (l *Lexer) ScanLine(source string) ([]Token, error) {
	l.Line += 1
	l.source = source
	l.curr = 0
	l.tokens = []Token{}

	for {
		l.start = l.curr
		if l.curr >= int64(len(source)) {
			break
		}
		err := l.scanToken()
		if err != nil {
			return nil, err
		}
	}

	return l.tokens, nil
}

func (l *Lexer) scanToken() error {
	var err error
	ch := l.advance()

	if ch == ' ' || ch == '\r' || ch == '\t' {
		return nil
	} else if ch == '\n' {
		l.Line += 1
	} else if ch == ';' {
		l.addToken(SEMICOLON)
	} else if ch == '{' {
		l.addToken(LEFT_BRACE)
	} else if ch == '}' {
		l.addToken(RIGHT_BRACE)
	} else if ch == '(' {
		l.addToken(LEFT_PAREN)
	} else if ch == ')' {
		l.addToken(RIGHT_PAREN)
	} else if ch == '+' {
		l.addToken(PLUS)
	} else if ch == '-' {
		l.addToken(MINUS)
	} else if ch == '*' {
		l.addToken(STAR)
	} else if ch == '!' {
		next := l.match('=')
		if next {
			l.addToken(BANG_EQ)
		} else {
			l.addToken(BANG)
		}
	} else if ch == '<' {
		next := l.match('=')
		if next {
			l.addToken(LESS_EQ)
		} else {
			l.addToken(LESS)
		}
	} else if ch == '>' {
		next := l.match('=')
		if next {
			l.addToken(GREATER_EQ)
		} else {
			l.addToken(GREATER)
		}
	} else if ch == '=' {
		next := l.match('=')
		if next {
			l.addToken(EQUAL_EQUAL)
		} else {
			l.addToken(EQUAL)
		}
	} else if ch == '/' {
		next := l.match('/')
		if next {
			for !l.atEnd() && l.peek() != '\n' {
				l.advance()
			}
			if l.peek() == '\n' {
				l.advance()
			}
			return nil
		} else {
			l.addToken(SLASH)
		}
	} else if ch == '"' {
		err = l.handleString()
		if err != nil {
			return err
		}
	} else if isNumber(ch) {
		if err = l.handleNumber(); err != nil {
			return err
		}
	} else if isAlpha(ch) {
		l.handleIdentifier()
	}
	return nil
}

func (l *Lexer) addToken(tokenType TokenType) {
	ch := l.source[l.start:l.curr]
	token := Token{
		Type:    tokenType,
		Literal: ch,
		Lexeme:  ch,
		Line:    l.Line,
	}
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) addTokenWithLiteral(tokenType TokenType, literal any) {
	ch := l.source[l.start:l.curr]
	token := Token{
		Type:    tokenType,
		Literal: literal,
		Lexeme:  ch,
		Line:    l.Line,
	}
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) handleIdentifier() {
	for {
		next := l.peek()
		if next == 0 || next == ' ' || !isAlphaNumeric(next) {
			break
		}
		if isAlpha(next) {
			l.advance()
		}
	}
	k := l.source[l.start:l.curr]
	tokenType, ok := keywords[k]
	if !ok {
		tokenType = IDENTIFIER
	}
	l.addToken(tokenType)
}

func (l *Lexer) handleNumber() error {
	for {
		next := l.peek()
		if isAlpha(next) {
			return errors.New("invalid number: contains alphabetic characters")
		}
		if next == 0 || next == ' ' || !isNumber(next) {
			break
		}
		if isNumber(next) {
			l.advance()
		}
	}
	literal, _ := Number(l.source[l.start:l.curr]) // todo: handle error
	l.addTokenWithLiteral(NUMBER, literal)
	return nil
}

func (l *Lexer) handleString() error {
	for {
		next := l.peek()
		if next == 0 || next == '"' {
			break
		}
		if next == '\n' {
			l.Line += 1
		}
		l.advance()
	}
	if l.atEnd() {
		return errors.New("unterminated string")
	}
	l.advance() // Skips the closing `"`
	l.addTokenWithLiteral(STRING, l.source[l.start+1:l.curr-1])
	return nil
}

func (l *Lexer) advance() byte {
	c := l.source[l.curr]
	l.curr += 1
	return c
}

func (l *Lexer) match(next byte) (matches bool) {
	if l.atEnd() {
		return false
	}
	if l.source[l.curr] != next {
		return false
	}
	l.curr += 1
	return true
}

func (l *Lexer) peek() (next byte) {
	if l.atEnd() {
		return 0
	}
	return l.source[l.curr]
}

func (l *Lexer) atEnd() bool {
	return l.curr >= int64(len(l.source))
}
