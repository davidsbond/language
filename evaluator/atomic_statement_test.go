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

func TestEvaluator_AtomicStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedKey    string
		ExpectedObject object.Object
	}{
		{
			Name:           "It should evaluate atomic number declarations",
			Expression:     "atomic test = 1",
			ExpectedKey:    "test",
			ExpectedObject: object.MakeAtomic(&object.Number{Value: 1}),
		},
		{
			Name:           "It should evaluate atomic string declarations",
			Expression:     `atomic test = "test"`,
			ExpectedKey:    "test",
			ExpectedObject: object.MakeAtomic(&object.String{Value: "test"}),
		},
		{
			Name:           "It should evaluate atomic bool declarations",
			Expression:     `atomic test = true`,
			ExpectedKey:    "test",
			ExpectedObject: object.MakeAtomic(&object.Boolean{Value: true}),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			ast := parser.Parse()

			scope := object.NewScope()
			evaluator.Evaluate(ast, scope)

			actual := scope.Get(tc.ExpectedKey)

			assert.NotNil(t, actual)
			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
