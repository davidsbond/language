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
	`

	rd := bufio.NewReader(strings.NewReader(code))
	lex, _ := lexer.New(rd)
	parser := parser.New(lex)
	ast, _ := parser.Parse()

	scope := object.NewScope()
	evaluator.Evaluate(ast, scope)
}
