// if yes, then assign the new value to that variable int the environment
// if yes, then assign the new value to that variable int the environment
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

func (e *Environment) Assign(name lexer.Token, value any) error {
	if _, ok := e.Values[name.Lexeme]; !ok {
		return fmt.Errorf("undefined variable: '%s'", name.Lexeme)
	}
	e.Values[name.Lexeme] = value
	return nil
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
