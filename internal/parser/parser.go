package parser

import "github.com/maxcelant/kiwi/internal/lexer"

type Parser struct {
	tokens  []lexer.Token
	current int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) parse() ([]Expr, error) {
	return nil, nil
}
