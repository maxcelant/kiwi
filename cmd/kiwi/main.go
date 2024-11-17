package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/maxcelant/kiwi/internal/lexer"
)

var lxr *lexer.Lexer

func run() error {
	var file *os.File
	var err error
	file, err = os.Open("test/sample.kiwi")
	if err != nil {
		return err
	}
	defer file.Close()

	lxr = lexer.New()
	tokens := []lexer.Token{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineTokens, err := lxr.ScanLine(scanner.Text())
		if err != nil {
			return err
		}
		tokens = append(tokens, lineTokens...)
	}

	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Lexeme: "", Literal: nil, Line: lxr.Line})
	fmt.Println(tokens)

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
