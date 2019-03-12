package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_AwaitStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name: "It should evaluate awaited async function calls",
			Expression: `
			async func add(a, b) { 
				return a + b 
			}
			
			await add(1, 2)`,
			ExpectedObject: &object.Number{
				Value: 3,
			},
		},
		{
			Name: "It should evaluate non-awaited async function calls",
			Expression: `
			async func add(a, b) { 
				return a + b 
			}
			
			add(1, 2)`,
			ExpectedObject: &object.Null{},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
