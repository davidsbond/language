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

func TestEvaluator_CallExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name: "It should evaluate valid call expressions",
			Expression: `
			func add(a, b) {
				return a + b
			}

			add(1, 2)
			`,
			ExpectedObject: &object.Number{
				Value: 3,
			},
		},
	}

	for _, tc := range tt {
		rd := bufio.NewReader(strings.NewReader(tc.Expression))
		lex, _ := lexer.New(rd)
		parser := parser.New(lex)
		ast, _ := parser.Parse()

		scope := object.NewScope()
		actual := evaluator.Evaluate(ast, scope)

		assert.NotNil(t, actual)
		assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
		assert.Equal(t, tc.ExpectedObject.String(), actual.String())
	}
}
