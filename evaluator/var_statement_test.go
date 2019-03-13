package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_VarStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate variable number declarations",
			Expression:     "var test = 1",
			ExpectedObject: &object.Number{Value: 1},
		},
		{
			Name:           "It should evaluate variable string declarations",
			Expression:     `var test = "test"`,
			ExpectedObject: &object.String{Value: "test"},
		},
		{
			Name:           "It should evaluate variable bool declarations",
			Expression:     "var test = true",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate variable character declarations",
			Expression:     "var test = 'a'",
			ExpectedObject: &object.Character{Value: 'a'},
		},
		{
			Name:       "It should evaluate constant array declarations",
			Expression: "var test = [1, 2, 3, 4]",
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
