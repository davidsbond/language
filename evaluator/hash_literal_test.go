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

func TestEvaluator_HashLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
		{
			Name:       "It should evaluate hash literals",
			Expression: `{ "a": 1, "b": "test", "c": 't' }`,
			ExpectedObject: &object.Hash{
				Pairs: map[object.HashKey]object.HashPair{
					object.HashKey{Type: object.TypeString, Value: 1}: object.HashPair{
						Key: &object.String{
							Value: "a",
						},
						Value: &object.Number{
							Value: 1,
						},
					},
					object.HashKey{Type: object.TypeString, Value: 2}: object.HashPair{
						Key: &object.String{
							Value: "b",
						},
						Value: &object.String{
							Value: "test",
						},
					},
					object.HashKey{Type: object.TypeString, Value: 3}: object.HashPair{
						Key: &object.String{
							Value: "c",
						},
						Value: &object.Character{
							Value: 't',
						},
					},
				},
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
			actual := evaluator.Evaluate(ast, scope)

			assert.NotNil(t, actual)

			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
