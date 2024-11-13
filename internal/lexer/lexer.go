package lexer

import (
	"errors"
	"strconv"
)

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

func isNumber(b byte) bool {
	return (b >= '0' && b <= '9')
}

func String(ch byte) string {
	return string(ch)
}

func Number(s string) (int, error) {
	return strconv.Atoi(s)
}

func (l *Lexer) ScanLine(line string) ([]Token, error) {
	l.line = line
	l.currLine += 1
	l.curr = 0
	l.tokens = []Token{}

	for {
		l.start = l.curr
		if l.curr >= int64(len(line)) {
			break
		}
		err := l.scanToken()
		if err != nil {
			return nil, err
		}
	}

	l.tokens = append(l.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: l.currLine})
	return l.tokens, nil
}

func (l *Lexer) scanToken() error {
	var token Token
	var err error
	ch := l.line[l.curr]

	if ch == ' ' || ch == '\r' || ch == '\t' {
		l.advance()
		return nil
	} else if ch == ';' {
		token = l.addToken(SEMICOLON)
	} else if ch == '!' {
		next := l.match('=')
		if next {
			l.advance()
			token = l.addToken(BANG_EQ)
		} else {
			token = l.addToken(BANG)
		}
	} else if ch == '<' {
		next := l.match('=')
		if next {
			l.advance()
			token = l.addToken(LESS_EQ)
		} else {
			token = l.addToken(LESS)
		}
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
	} else if isNumber(ch) {
		token, err = l.handleNumber()
		if err != nil {
			return err
		}
	}
	l.tokens = append(l.tokens, token)
	l.advance()
	return nil
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

func (l *Lexer) handleNumber() (Token, error) {
	for {
		next := l.peek()
		if next == 0 || next == ' ' {
			break
		}
		if isAlpha(next) {
			return Token{}, errors.New("invalid number: contains alphabetic characters")
		}
		if isNumber(next) {
			l.advance()
		}
	}
	literal, _ := Number(l.line[l.start : l.curr+1]) // todo: handle error
	return l.addTokenWithLiteral(NUMBER, literal), nil
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
	l.advance() // Skips the closing `"`
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
