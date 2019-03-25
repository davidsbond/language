package evaluator_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func TestEvaluator_CharacterInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:           "It should evaluate addition",
			Expression:     `'a' + 'b'`,
			ExpectedObject: &object.String{Value: "ab"},
		},
		{
			Name:           "It should evaluate subtraction",
			Expression:     `'b' - 'a'`,
			ExpectedObject: &object.Character{Value: 1},
		},
		{
			Name:           "It should evaluate multiplication",
			Expression:     `'a' * 'a'`,
			ExpectedObject: &object.Character{Value: 'Ⓛ'},
		},
		{
			Name:           "It should evaluate division",
			Expression:     `'a' / 'b'`,
			ExpectedObject: &object.Character{Value: 0},
		},
		{
			Name:           "It should evaluate greater than",
			Expression:     `'b' > 'a'`,
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate less than",
			Expression:     `'a' < 'b'`,
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate equality",
			Expression:     `'a' == 'a'`,
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should addition of char and string",
			Expression:     `'a' + "a"`,
			ExpectedObject: &object.String{Value: "aa"},
		},
		{
			Name:           "It should evaluate less than or equal to",
			Expression:     "'a' <= 'b'",
			ExpectedObject: &object.Boolean{Value: true},
		},
		{
			Name:           "It should evaluate greater than or equal to",
			Expression:     "'b' >= 'a'",
			ExpectedObject: &object.Boolean{Value: true},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
