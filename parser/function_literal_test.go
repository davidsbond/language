package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_FunctionLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name: "It should parse valid inline function literals",
			Expression: `
			var add = func(a, b) {
				return a + b
			}
			`,
			ExpectedNode: &ast.VarStatement{
				Token: token.New(token.VAR, token.VAR, 0, 0),
				Name: &ast.Identifier{
					Token: token.New("add", token.VAR, 0, 0),
					Value: "add",
				},
				Value: &ast.FunctionLiteral{
					Name: &ast.Identifier{
						Token: token.New("func", token.IDENT, 0, 0),
						Value: "anonymous",
					},
					Parameters: []*ast.Identifier{
						{
							Token: token.New("a", token.IDENT, 0, 0),
							Value: "a",
						},
						{
							Token: token.New("b", token.IDENT, 0, 0),
							Value: "b",
						},
					},
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
		},
		{
			Name: "It should parse a valid function literal",
			Expression: `
			func add(a, b) {
				return a + b
			}`,
			ExpectedNode: &ast.FunctionLiteral{
				Name: &ast.Identifier{
					Token: token.New("add", token.IDENT, 0, 0),
					Value: "add",
				},
				Parameters: []*ast.Identifier{
					{
						Token: token.New("a", token.IDENT, 0, 0),
						Value: "a",
					},
					{
						Token: token.New("b", token.IDENT, 0, 0),
						Value: "b",
					},
				},
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
