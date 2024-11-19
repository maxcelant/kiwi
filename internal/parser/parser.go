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
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) parse() (Expr, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, fmt.Errorf("parsing error occurred: %w", err)
	}
	return expr, nil
}

func (p *Parser) expression() (Expr, error) {

}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Type == lexer.EOF
}
