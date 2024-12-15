package env

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/lexer"
)

type Environment struct {
	Values map[string]any
}

func New() Environment {
	return Environment{
		Values: make(map[string]any),
	}
}

func (e *Environment) Define(name string, value any) {
	e.Values[name] = value
}

func (e *Environment) Get(token lexer.Token) (any, error) {
	value, ok := e.Values[token.Lexeme]
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", token.Lexeme)
	}
	return value, nil
}
