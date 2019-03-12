package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_CallExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse a valid call expression",
			Expression: "test(1, 2, 3)",
			ExpectedNode: &ast.CallExpression{
				Function: &ast.Identifier{
					Value: "test",
					Token: token.New("test", token.IDENT, 0, 0),
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
					&ast.NumberLiteral{
						Token: token.New("3", token.NUMBER, 0, 0),
						Value: 3,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
