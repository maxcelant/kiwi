package main

import (
	"fmt"
	"log"
	"os"

	"github.com/maxcelant/kiwi/internal/interpreter"
	"github.com/maxcelant/kiwi/internal/lexer"
	"github.com/maxcelant/kiwi/internal/parser"
)

var lxr *lexer.Lexer
var psr *parser.Parser
var it *interpreter.Interpreter

// todo: fix me
// func runREPL() error {
// 	var file *os.File
// 	var err error
// 	file, err = os.Open("test/sample.kiwi")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	lxr = lexer.New()
// 	tokens := []lexer.Token{}
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		lineTokens, err := lxr.ScanLine(scanner.Text())
// 		if err != nil {
// 			return err
// 		}
// 		tokens = append(tokens, lineTokens...)
// 	}

// 	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Lexeme: "", Literal: nil, Line: lxr.Line})
// 	fmt.Println(tokens)

// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}
// 	return nil
// }

func run() error {
	content, err := os.ReadFile("test/sample.kiwi")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	lxr = lexer.New(string(content))
	tokens, err := lxr.Scan()
	if err != nil {
		return fmt.Errorf("syntax error occurred: %w", err)
	}

	psr = parser.New(tokens)
	expr, err := psr.Parse()
	if err != nil {
		return fmt.Errorf("parse error occurred: %w", err)
	}

	it = interpreter.New(expr)
	it.Interpret()

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	// if err := runREPL(); err != nil {
	// 	log.Fatal(err)
	// }
}
