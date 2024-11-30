package parser

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/lexer"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() (Expr, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}
	expr, err := p.expression()
	if err != nil {
		return nil, fmt.Errorf("parsing error occurred: %w", err)
	}
	return expr, nil
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	return expr, err
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()

	for p.match(lexer.LESS, lexer.LESS_EQ, lexer.GREATER, lexer.GREATER_EQ) {
		operator := p.prev()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &Binary{
			right:    right,
			operator: operator,
			left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()

	for p.match(lexer.PLUS, lexer.MINUS) {
		operator := p.prev()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = &Binary{
			right:    right,
			operator: operator,
			left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()

	for p.match(lexer.SLASH, lexer.STAR) {
		operator := p.prev()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &Binary{
			right:    right,
			operator: operator,
			left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) unary() (Expr, error) {
	if p.match(lexer.BANG, lexer.MINUS) {
		operator := p.prev()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &Unary{
			operator: operator,
			right:    right,
		}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(lexer.TRUE) {
		return &Primary{value: true}, nil
	}
	if p.match(lexer.FALSE) {
		return &Primary{value: false}, nil
	}
	if p.match(lexer.NIL) {
		return &Primary{value: nil}, nil
	}
	if p.match(lexer.STRING) {
		return &Primary{value: p.prev().Literal}, nil
	}
	if p.match(lexer.NUMBER) {
		return &Primary{value: p.prev().Literal}, nil
	}

	return nil, fmt.Errorf("%s expected expression", p.peek().Lexeme)
}

func (p *Parser) match(matchers ...lexer.TokenType) bool {
	for _, m := range matchers {
		if p.check(m) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.prev()
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) prev() lexer.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Type == lexer.EOF
}
