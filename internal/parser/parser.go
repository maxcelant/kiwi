package parser

import (
	"errors"
	"fmt"

	exp "github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() (exp.Expr, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}
	expr, err := p.expression()
	if err != nil {
		return nil, fmt.Errorf("parsing error occurred: %w", err)
	}
	return expr, nil
}

func (p *Parser) expression() (exp.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (exp.Expr, error) {
	expr, err := p.comparison()

	for p.match(lexer.EQUAL, lexer.BANG_EQ) {
		operator := p.prev()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = exp.Binary{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) comparison() (exp.Expr, error) {
	expr, err := p.term()

	for p.match(lexer.LESS, lexer.LESS_EQ, lexer.GREATER, lexer.GREATER_EQ) {
		operator := p.prev()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = exp.Binary{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) term() (exp.Expr, error) {
	expr, err := p.factor()

	for p.match(lexer.PLUS, lexer.MINUS) {
		operator := p.prev()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = exp.Binary{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) factor() (exp.Expr, error) {
	expr, err := p.unary()

	for p.match(lexer.SLASH, lexer.STAR) {
		operator := p.prev()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = exp.Binary{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) unary() (exp.Expr, error) {
	if p.match(lexer.BANG, lexer.MINUS) {
		operator := p.prev()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return exp.Unary{
			Operator: operator,
			Right:    right,
		}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (exp.Expr, error) {
	if p.match(lexer.TRUE) {
		return exp.Primary{Value: true}, nil
	}
	if p.match(lexer.FALSE) {
		return exp.Primary{Value: false}, nil
	}
	if p.match(lexer.NIL) {
		return exp.Primary{Value: nil}, nil
	}
	if p.match(lexer.STRING) {
		return exp.Primary{Value: p.prev().Literal}, nil
	}
	if p.match(lexer.NUMBER) {
		return exp.Primary{Value: p.prev().Literal}, nil
	}
	if p.match(lexer.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		err = p.consume(lexer.RIGHT_PAREN, "Expected right parent ')' after expression")
		if err != nil {
			return nil, err
		}
		return exp.Grouping{Expr: expr}, nil
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

func (p *Parser) consume(tokenType lexer.TokenType, err string) error {
	if p.peek().Type == tokenType {
		p.advance()
		return nil
	}
	return errors.New(err)
}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Type == lexer.EOF
}
