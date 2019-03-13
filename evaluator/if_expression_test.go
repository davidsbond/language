package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/evaluator"
)

func TestEvaluator_IfExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate if expression consequence",
			Expression:     "if (true) { return false } else { return true }",
			ExpectedObject: evaluator.FALSE,
		},
		{
			Name:           "It should evaluate if expression alternative",
			Expression:     "if (false) { return false } else { return true }",
			ExpectedObject: evaluator.TRUE,
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
