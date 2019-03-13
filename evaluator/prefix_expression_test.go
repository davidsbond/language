package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_PrefixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate minus prefixes",
			Expression:     "-1",
			ExpectedObject: &object.Number{Value: -1},
		},
		{
			Name:           "It should evaluate bang prefixes",
			Expression:     "!true",
			ExpectedObject: &object.Boolean{Value: false},
		},
		{
			Name:           "It should evaluate sqrt prefixes",
			Expression:     "√4",
			ExpectedObject: &object.Number{Value: 2},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
