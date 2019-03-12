package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_ArrayLiteral(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:       "It should evaluate array literals",
			Expression: "[1, 2, 3, 4]",
			ExpectedObject: &object.Array{
				Elements: []object.Object{
					&object.Number{Value: 1},
					&object.Number{Value: 2},
					&object.Number{Value: 3},
					&object.Number{Value: 4},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
