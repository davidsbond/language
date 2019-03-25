package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func TestEvaluator_Identifier(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name: "It should evaluate an identifier",
			Expression: `
			var a = 1 + 1

			a
			`,
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name: "It should evaluate a built-in call",
			Expression: `
			var a = 1 + 1

			type(a)
			`,
			ExpectedObject: &object.String{Value: object.TypeNumber},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
