package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func TestEvaluator_ReturnStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:       "It should evaluate string return statements",
			Expression: `return "test"`,
			ExpectedObject: &object.String{
				Value: "test",
			},
		},
		{
			Name:       "It should evaluate number return statements",
			Expression: `return 1`,
			ExpectedObject: &object.Number{
				Value: 1,
			},
		},
		{
			Name:       "It should evaluate character return statements",
			Expression: `return 'a'`,
			ExpectedObject: &object.Character{
				Value: 'a',
			},
		},
		{
			Name:       "It should evaluate boolean return statements",
			Expression: `return true`,
			ExpectedObject: &object.Boolean{
				Value: true,
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
