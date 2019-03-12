package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/token"

	"github.com/davidsbond/dave/ast"
)

func TestParser_AwaitStatement(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name: "It should parse valid await statements",
			Expression: `
			await add(1, 2)
			`,
			ExpectedNode: &ast.AwaitStatement{
				Token: token.New("await", token.AWAIT, 0, 0),
				Value: &ast.CallExpression{
					Function: &ast.Identifier{
						Value: "add",
						Token: token.New("add", token.IDENT, 0, 0),
					},
					Arguments: []ast.Node{
						&ast.NumberLiteral{
							Token: token.New("1", token.NUMBER, 0, 0),
							Value: 1,
						},
						&ast.NumberLiteral{
							Token: token.New("2", token.NUMBER, 0, 0),
							Value: 2,
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
