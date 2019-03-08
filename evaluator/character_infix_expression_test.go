package evaluator_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/evaluator"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/parser"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator_CharacterInfixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}{
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
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			ast := parser.Parse()

			scope := object.NewScope()
			actual := evaluator.Evaluate(ast, scope)

			assert.NotNil(t, actual)
			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
