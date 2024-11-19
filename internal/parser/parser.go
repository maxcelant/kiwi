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

func (p *Parser) parse() (Expr, error) {
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
	expr := p.comparison()

}

func (p *Parser) comparison() (Expr, error) {

}

func (p *Parser) term() (Expr, error) {

}

func (p *Parser) factor() (Expr, error) {

}

func (p *Parser) unary() (Expr, error) {

}

func (p *Parser) primary() (Expr, error) {

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
