package main

import (
	"fmt"
	"log"
	"os"

	"github.com/maxcelant/kiwi/internal/lexer"
)

var lxr *lexer.Lexer

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

func runFile() error {
	content, err := os.ReadFile("test/sample.kiwi")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	lxr = lexer.New(string(content))
	tokens, err := lxr.Scan()
	if err != nil {
		return fmt.Errorf("syntax error occurred: %w", err)
	}

	fmt.Println(tokens)

	return nil
}

func main() {
	if err := runFile(); err != nil {
		log.Fatal(err)
	}
	// if err := runREPL(); err != nil {
	// 	log.Fatal(err)
	// }
}
