package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
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
		{
			Name:           "It should addition of string and char",
			Expression:     `"a" + 'a'`,
			ExpectedObject: &object.String{Value: "aa"},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
