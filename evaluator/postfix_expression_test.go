package evaluator_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/evaluator"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/parser"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator_PostfixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name: "It should evaluate decremental postfixes",
			Expression: `
			var a = 1
			a--
			`,
			ExpectedObject: &object.Number{Value: 0},
		},
		{
			Name: "It should evaluate incremental postfixes",
			Expression: `
			var a = 1
			a++
			`,
			ExpectedObject: &object.Number{Value: 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			ast, _ := parser.Parse()

			scope := object.NewScope()
			actual := evaluator.Evaluate(ast, scope)

			assert.NotNil(t, actual)
			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
