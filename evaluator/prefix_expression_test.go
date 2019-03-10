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

func TestEvaluator_PrefixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name:           "It should evaluate minus prefixes",
			Expression:     "-1",
			ExpectedObject: &object.Number{Value: -1},
		},
		{
			Name:           "It should evaluate bang prefixes",
			Expression:     "!true",
			ExpectedObject: &object.Boolean{Value: false},
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
