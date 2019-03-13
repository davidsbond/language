package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_BooleanInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate equality",
			Expression:     `true == true`,
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate not equals",
			Expression:     `true != false`,
			ExpectedObject: &object.Boolean{Value: true},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
