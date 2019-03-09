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
	async func testAsync() {
		return 1 + 1
	}

	const result = await testAsync()

	testAsync()
	`

	rd := bufio.NewReader(strings.NewReader(code))
	lex, _ := lexer.New(rd)
	parser := parser.New(lex)
	ast, _ := parser.Parse()

	scope := object.NewScope()
	evaluator.Evaluate(ast, scope)
}
