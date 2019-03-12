package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_AtomicStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate atomic number declarations",
			Expression:     "atomic test = 1",
			ExpectedObject: object.MakeAtomic(&object.Number{Value: 1}),
		},
		{
			Name:           "It should evaluate atomic string declarations",
			Expression:     `atomic test = "test"`,
			ExpectedObject: object.MakeAtomic(&object.String{Value: "test"}),
		},
		{
			Name:           "It should evaluate atomic bool declarations",
			Expression:     `atomic test = true`,
			ExpectedObject: object.MakeAtomic(&object.Boolean{Value: true}),
		},
		{
			Name:           "It should evaluate atomic character declarations",
			Expression:     "atomic test = 'a'",
			ExpectedObject: object.MakeAtomic(&object.Character{Value: 'a'}),
		},
		{
			Name:       "It should evaluate atomic array declarations",
			Expression: "atomic test = [1, 2, 3, 4]",
			ExpectedObject: object.MakeAtomic(&object.Array{
				Elements: []object.Object{
					&object.Number{Value: 1},
					&object.Number{Value: 2},
					&object.Number{Value: 3},
					&object.Number{Value: 4},
				}}),
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
