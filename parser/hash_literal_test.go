package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_HashLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name: "It should parse valid hash literals",
			Expression: `{ 
				"a": 1, 
				"b": "test", 
				"c": 't' 
			}`,
			ExpectedNode: &ast.HashLiteral{
				Token: token.New(token.LBRACE, token.LBRACE, 0, 0),
				Pairs: map[ast.Node]ast.Node{
					&ast.StringLiteral{
						Value: "a",
						Token: token.New("a", token.STRING, 0, 0),
					}: &ast.NumberLiteral{
						Value: 1,
						Token: token.New("1", token.NUMBER, 0, 0),
					},
					&ast.StringLiteral{
						Value: "b",
						Token: token.New("b", token.STRING, 0, 0),
					}: &ast.StringLiteral{
						Value: "test",
						Token: token.New("test", token.STRING, 0, 0),
					},
					&ast.StringLiteral{
						Value: "c",
						Token: token.New("c", token.STRING, 0, 0),
					}: &ast.CharacterLiteral{
						Value: 't',
						Token: token.New("t", token.CHAR, 0, 0),
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
