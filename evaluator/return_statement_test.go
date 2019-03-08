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

func TestEvaluator_ReturnStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name:       "It should evaluate string return statements",
			Expression: `return "test"`,
			ExpectedObject: &object.String{
				Value: "test",
			},
		},
		{
			Name:       "It should evaluate number return statements",
			Expression: `return 1`,
			ExpectedObject: &object.Number{
				Value: 1,
			},
		},
		{
			Name:       "It should evaluate character return statements",
			Expression: `return 'a'`,
			ExpectedObject: &object.Character{
				Value: 'a',
			},
		},
		{
			Name:       "It should evaluate boolean return statements",
			Expression: `return true`,
			ExpectedObject: &object.Boolean{
				Value: true,
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
