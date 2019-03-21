package evaluator_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func TestEvaluator_AsyncStatement(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name:       "It should evaluate async function declarations",
			Expression: "async func add(a, b) { return a + b }",
			ExpectedObject: &object.Function{
				Name: &ast.Identifier{
					Value: "add",
					Token: token.New("add", token.IDENT, 0, 0),
				},
				Parameters: []*ast.Identifier{
					{
						Value: "a",
						Token: token.New("a", token.IDENT, 0, 0),
					},
					{
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
		tc.Run(t)
	}
}
