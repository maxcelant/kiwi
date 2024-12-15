package parser

import (
	"errors"
	"fmt"

	exp "github.com/maxcelant/kiwi/internal/expr"
	"github.com/maxcelant/kiwi/internal/lexer"
	"github.com/maxcelant/kiwi/internal/stmt"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() ([]stmt.Stmt, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}
	statements := []stmt.Stmt{}
	for !p.isAtEnd() {
		s, err := p.declaration()
		if err != nil {
			return nil, fmt.Errorf("parsing error occurred: %w", err)
		}
		statements = append(statements, s)
	}
	return statements, nil
}

func (p *Parser) declaration() (stmt.Stmt, error) {
	if p.match(lexer.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) varDeclaration() (stmt.Stmt, error) {
	var err error
	var initializer exp.Expr
	name, err := p.consume(lexer.IDENTIFIER, "expect variable name")
	if err != nil {
		return nil, err
	}
	if p.match(lexer.EQUAL) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	_, err = p.consume(lexer.SEMICOLON, "expect ';' after value")
	if err != nil {
		return nil, err
	}
	return stmt.Var{
		Name:        name,
		Initializer: initializer,
	}, nil
}

func (p *Parser) statement() (stmt.Stmt, error) {
	if p.match(lexer.PRINT) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() (stmt.Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(lexer.SEMICOLON, "expect ';' after value")
	if err != nil {
		return nil, err
	}
	return stmt.Print{
		Expression: expr,
	}, nil
}

func (p *Parser) expressionStatement() (stmt.Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(lexer.SEMICOLON, "expect ';' after value")
	if err != nil {
		return nil, err
	}
	return stmt.Expression{
		Expression: expr,
	}, nil
}

func (p *Parser) expression() (exp.Expr, error) {
	return p.assignment()
}

func (p *Parser) assignment() (exp.Expr, error) {
	expr, err := p.logicOR()

	if p.match(lexer.EQUAL) {
		value, err := p.assignment()
		if err != nil {
			return nil, err
		}
		// This allows us to verify the left-hand value is an expression that *can*
		// have something assigned to it, otherwise, this in invalid
		variable, ok := expr.(exp.Variable)
		if !ok {
			return nil, fmt.Errorf("invalid assignment target")
		}
		name := variable.Name
		return exp.Assign{
			Name:  name,
			Value: value,
		}, nil
	}

	return expr, err
}

func (p *Parser) logicOR() (exp.Expr, error) {
	expr, err := p.logicAND()

	for p.match(lexer.OR) {
		operator := p.prev()
		right, err := p.logicAND()
		if err != nil {
			return nil, err
		}
		expr = exp.Logical{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) logicAND() (exp.Expr, error) {
	expr, err := p.equality()

	for p.match(lexer.AND) {
		operator := p.prev()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}
		expr = exp.Logical{
			Right:    right,
			Operator: operator,
			Left:     expr,
		}
	}

	return expr, err
}

func (p *Parser) equality() (exp.Expr, error) {
	expr, err := p.comparison()

	for p.match(lexer.EQUAL_EQUAL, lexer.BANG_EQ) {
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
	if p.match(lexer.IDENTIFIER) {
		return exp.Variable{Name: p.prev()}, nil
	}
	if p.match(lexer.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(lexer.RIGHT_PAREN, "Expected right parent ')' after expression")
		if err != nil {
			return nil, err
		}
		return exp.Grouping{Expression: expr}, nil
	}

	token, _ := p.peek()
	return nil, fmt.Errorf("%s expected expression", token.Lexeme)
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
	next, err := p.peek()
	if err != nil {
		return false
	}
	return next.Type == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.prev()
}

func (p *Parser) peek() (lexer.Token, error) {
	if p.isAtEnd() {
		return lexer.Token{}, errors.New("reached end of file")
	}
	return p.tokens[p.current], nil
}

func (p *Parser) consume(tokenType lexer.TokenType, errMsg string) (lexer.Token, error) {
	next, err := p.peek()
	if err != nil {
		return lexer.Token{}, fmt.Errorf("%s: %w", errMsg, err)
	}
	if next.Type == tokenType {
		return p.advance(), nil
	}
	return lexer.Token{}, errors.New(errMsg)
}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Type == lexer.EOF
}

func (p *Parser) prev() lexer.Token {
	return p.tokens[p.current-1]
}
