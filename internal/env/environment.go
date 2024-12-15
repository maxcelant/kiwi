package env

import (
	"fmt"

	"github.com/maxcelant/kiwi/internal/lexer"
)

// Holds a reference to the "enclosing/parent" environment so that when its scope ends,
// We can return to the encompassing scope.
type Environment struct {
	Values map[string]any
	Parent *Environment
}

func New(parent *Environment) *Environment {
	return &Environment{
		Values: make(map[string]any),
		Parent: parent,
	}
}

func (e *Environment) Assign(name lexer.Token, value any) error {
	if _, ok := e.Values[name.Lexeme]; ok {
		e.Values[name.Lexeme] = value
		return nil
	}

	if e.Parent != nil {
		return e.Parent.Assign(name, value)
	}

	return fmt.Errorf("undefined variable: '%s'", name.Lexeme)
}

func (e *Environment) Define(name string, value any) {
	e.Values[name] = value
}

func (e *Environment) Get(token lexer.Token) (any, error) {
	if value, ok := e.Values[token.Lexeme]; ok {
		return value, nil
	}

	if e.Parent != nil {
		return e.Parent.Get(token)
	}

	return nil, fmt.Errorf("undefined variable: %s", token.Lexeme)
}
