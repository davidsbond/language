package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_StringInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate addition",
			Expression:     `"a" + "b"`,
			ExpectedObject: &object.String{Value: "ab"},
		},
		{
			Name:           "It should evaluate equality",
			Expression:     `"a" == "a"`,
			ExpectedObject: &object.Boolean{Value: true},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
