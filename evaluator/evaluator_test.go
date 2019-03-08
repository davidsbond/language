package evaluator_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/evaluator"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/parser"
)

func TestEvaluator_Evaluate(t *testing.T) {
	code := `
	const a = "test"
	atomic b = "test" + a
	var c = "test" + b

	func add(a, b) {
		return a + b
	}

	var sub = func(a, b) {
		return a - b
	}

	const added = add(1, 2)
	const subbed = sub(2, 1)
	`

	rd := bufio.NewReader(strings.NewReader(code))
	lex, _ := lexer.New(rd)
	parser := parser.New(lex)
	ast, _ := parser.Parse()

	scope := object.NewScope()
	evaluator.Evaluate(ast, scope)
}
