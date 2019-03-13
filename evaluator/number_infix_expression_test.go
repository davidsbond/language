package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_NumberInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate addition",
			Expression:     "1 + 1",
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name:           "It should evaluate subtraction",
			Expression:     "1 - 1",
			ExpectedObject: &object.Number{Value: 0},
		},
		{
			Name:           "It should evaluate multiplication",
			Expression:     "1 * 2",
			ExpectedObject: &object.Number{Value: 2},
		},
		{
			Name:           "It should evaluate division",
			Expression:     "1 / 2",
			ExpectedObject: &object.Number{Value: 0.5},
		},
		{
			Name:           "It should evaluate modulo",
			Expression:     "1 % 2",
			ExpectedObject: &object.Number{Value: 1},
		},
		{
			Name:           "It should evaluate greater than",
			Expression:     "2 > 1",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate less than",
			Expression:     "1 < 2",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate equality",
			Expression:     "1 == 1",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name: "It should evaluate variable reassignment",
			Expression: `
			var a = 1
			a = 2
			`,
			ExpectedObject: &object.Number{Value: 2},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
