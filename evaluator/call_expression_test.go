package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_CallExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name: "It should evaluate valid, inline call expressions",
			Expression: `
			var add = func(a, b) {
				return a + b
			}

			add(1, 2)
			`,
			ExpectedObject: &object.Number{
				Value: 3,
			},
		},
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
		tc.Run(t)
	}
}
