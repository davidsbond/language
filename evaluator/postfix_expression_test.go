package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func TestEvaluator_PostfixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
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
		tc.Run(t)
	}
}
