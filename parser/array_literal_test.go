package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/token"

	"github.com/davidsbond/dave/ast"
)

func TestParser_ArrayLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse valid array literals",
			Expression: `[1, "test", 't']`,
			ExpectedNode: &ast.ArrayLiteral{
				Token: token.New(token.TRUE, token.TRUE, 0, 0),
				Elements: []ast.Node{
					&ast.NumberLiteral{
						Token: token.New("1", token.NUMBER, 0, 0),
						Value: 1,
					},
					&ast.StringLiteral{
						Token: token.New("test", token.STRING, 0, 0),
						Value: "test",
					},
					&ast.CharacterLiteral{
						Token: token.New("t", token.CHAR, 0, 0),
						Value: 't',
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
