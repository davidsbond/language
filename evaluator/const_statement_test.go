package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func TestEvaluator_ConstStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:       "It should evaluate constant number declarations",
			Expression: "const test = 1",
			ExpectedObject: &object.Constant{
				Value: &object.Number{
					Value: 1,
				},
			},
		},
		{
			Name:       "It should evaluate constant string declarations",
			Expression: `const test = "test"`,
			ExpectedObject: &object.Constant{
				Value: &object.String{
					Value: "test",
				},
			},
		},
		{
			Name:       "It should evaluate constant bool declarations",
			Expression: `const test = true`,
			ExpectedObject: &object.Constant{
				Value: &object.Boolean{
					Value: true,
				},
			},
		},
		{
			Name:       "It should evaluate constant character declarations",
			Expression: "const test = 'a'",
			ExpectedObject: &object.Constant{
				Value: &object.Character{Value: 'a'},
			},
		},
		{
			Name:       "It should evaluate constant array declarations",
			Expression: "const test = [1, 2, 3, 4]",
			ExpectedObject: &object.Constant{
				Value: &object.Array{
					Elements: []object.Object{
						&object.Number{Value: 1},
						&object.Number{Value: 2},
						&object.Number{Value: 3},
						&object.Number{Value: 4},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
