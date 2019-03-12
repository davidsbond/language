package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_IndexExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
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
		tc.Run(t)
	}
}
