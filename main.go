package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/davidsbond/language/evaluator"
	"github.com/davidsbond/language/lexer"
	"github.com/davidsbond/language/object"
	"github.com/davidsbond/language/parser"
)

func main() {
	scope := object.NewScope()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

repl:
	for scanner.Scan() {
		line := bytes.NewBuffer(scanner.Bytes())

		lexer, _ := lexer.New(bufio.NewReader(line))
		parser := parser.New(lexer)
		ast, errs := parser.Parse()

		for _, err := range errs {
			fmt.Println(err.Error())
			fmt.Print("> ")
			line.Reset()

			continue repl
		}

		if result := evaluator.Evaluate(ast, scope); result != nil {
			fmt.Println(result.String())
		}

		fmt.Print("> ")
	}
}
