package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/object"
)

func TestEvaluator_HashLiteral(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:       "It should evaluate hash literals",
			Expression: `{ "a": 1, "b": "test", "c": 't' }`,
			ExpectedObject: &object.Hash{
				Pairs: map[object.HashKey]object.HashPair{
					object.HashKey{Type: object.TypeString, Value: 12638187200555643000}: {
						Key: &object.String{
							Value: "a",
						},
						Value: &object.Number{
							Value: 1,
						},
					},
					object.HashKey{Type: object.TypeString, Value: 12638190499090526000}: {
						Key: &object.String{
							Value: "b",
						},
						Value: &object.String{
							Value: "test",
						},
					},
					object.HashKey{Type: object.TypeString, Value: 12638189399578898000}: {
						Key: &object.String{
							Value: "c",
						},
						Value: &object.Character{
							Value: 't',
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
