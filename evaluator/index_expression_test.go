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

func TestEvaluator_IndexExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name: "It should evaluate array index expressions",
			Expression: `
			const a = [1, 2, 3]
			var b = a[1]
			`,
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name: "It should evaluate map index expressions",
			Expression: `
			const a = { "test": "test" }
			var b = a["test"]
			`,
			ExpectedObject: &object.String{Value: "test"},
		},
		{
			Name: "It should evaluate string index expressions",
			Expression: `
			const a = "test"
			var b = a[1]
			`,
			ExpectedObject: &object.Character{Value: 'e'},
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
