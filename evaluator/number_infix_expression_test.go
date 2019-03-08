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

func TestEvaluator_NumberInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name:           "It should evaluate addition",
			Expression:     "1 + 1",
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name:           "It should evaluate subtraction",
			Expression:     "1 - 1",
			ExpectedObject: &object.Number{Value: 0},
		},
		{
			Name:           "It should evaluate multiplication",
			Expression:     "1 * 2",
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name:           "It should evaluate division",
			Expression:     "1 / 2",
			ExpectedObject: &object.Number{Value: 0.5},
		},
		{
			Name:           "It should evaluate modulo",
			Expression:     "1 % 2",
			ExpectedObject: &object.Number{Value: 1},
		},
		{
			Name:           "It should evaluate greater than",
			Expression:     "2 > 1",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate less than",
			Expression:     "1 < 2",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate equality",
			Expression:     "1 == 1",
			ExpectedObject: &object.Boolean{Value: true},
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
