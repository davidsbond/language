package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_ReturnStatement(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse string literals",
			Expression: `return "test"`,
			ExpectedNode: &ast.ReturnStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				ReturnValue: &ast.StringLiteral{
					Token: token.New(`"`, token.STRING, 0, 0),
					Value: "test",
				},
			},
		},
		{
			Name:       "It should parse numbers",
			Expression: `return 1`,
			ExpectedNode: &ast.ReturnStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				ReturnValue: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse characters",
			Expression: `return '1'`,
			ExpectedNode: &ast.ReturnStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				ReturnValue: &ast.CharacterLiteral{
					Token: token.New("'", token.CHAR, 0, 0),
					Value: '1',
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
