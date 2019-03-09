package evaluator_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/evaluator"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/parser"
	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator_AsyncStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Expression     string
		ExpectedKey    string
		ExpectedObject object.Object
	}{
		{
			Name:        "It should evaluate async function declarations",
			Expression:  "async func add(a, b) { return a + b }",
			ExpectedKey: "add",
			ExpectedObject: &object.Function{
				Name: &ast.Identifier{
					Value: "add",
					Token: token.New("add", token.IDENT, 0, 0),
				},
				Parameters: []*ast.Identifier{
					&ast.Identifier{
						Value: "a",
						Token: token.New("a", token.IDENT, 0, 0),
					},
					&ast.Identifier{
						Value: "b",
						Token: token.New("b", token.IDENT, 0, 0),
					},
				},
				Async: true,
				Body: &ast.BlockStatement{
					Token: token.New("{", token.LBRACE, 0, 0),
					Statements: []ast.Node{
						&ast.ReturnStatement{
							Token: token.New(token.RETURN, token.RETURN, 0, 0),
							ReturnValue: &ast.InfixExpression{
								Token:    token.New(token.PLUS, token.PLUS, 0, 0),
								Operator: "+",
								Left: &ast.Identifier{
									Token: token.New("a", token.IDENT, 0, 0),
									Value: "a",
								},
								Right: &ast.Identifier{
									Token: token.New("b", token.IDENT, 0, 0),
									Value: "b",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			ast, _ := parser.Parse()

			scope := object.NewScope()
			evaluator.Evaluate(ast, scope)

			actual := scope.Get(tc.ExpectedKey)

			assert.NotNil(t, actual)

			assert.Equal(t, tc.ExpectedObject.Type(), actual.Type())
			assert.Equal(t, tc.ExpectedObject.String(), actual.String())
		})
	}
}
