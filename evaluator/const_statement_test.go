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

func TestEvaluator_ConstStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Expression     string
		ExpectedKey    string
		Name           string
		ExpectedObject object.Object
	}{
		{
			Name:        "It should evaluate constant number declarations",
			Expression:  "const test = 1",
			ExpectedKey: "test",
			ExpectedObject: &object.Constant{
				Value: &object.Number{
					Value: 1,
				},
			},
		},
		{
			Name:        "It should evaluate constant string declarations",
			Expression:  `const test = "test"`,
			ExpectedKey: "test",
			ExpectedObject: &object.Constant{
				Value: &object.String{
					Value: "test",
				},
			},
		},
		{
			Name:        "It should evaluate constant bool declarations",
			Expression:  `const test = true`,
			ExpectedKey: "test",
			ExpectedObject: &object.Constant{
				Value: &object.Boolean{
					Value: true,
				},
			},
		},
		{
			Name:        "It should evaluate constant character declarations",
			Expression:  "var test = 'a'",
			ExpectedKey: "test",
			ExpectedObject: &object.Constant{
				Value: &object.Character{Value: 'a'},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			ast, _ := parser.Parse()

			scope := object.NewScope()
			evaluator.Evaluate(ast, scope)

			actual := scope.Get(tc.ExpectedKey)

			assert.NotNil(t, actual)
			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
