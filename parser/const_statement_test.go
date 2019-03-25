package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_ConstStatement(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse constant number declarations",
			Expression: "const test = 1",
			ExpectedNode: &ast.ConstStatement{
				Token: token.New(token.CONST, token.CONST, 0, 0),
				Name: &ast.Identifier{
					Token: token.New("test", token.IDENT, 0, 0),
					Value: "test",
				},
				Value: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse constant string declarations",
			Expression: `const test = "test"`,
			ExpectedNode: &ast.ConstStatement{
				Token: token.New(token.CONST, token.CONST, 0, 0),
				Name: &ast.Identifier{
					Token: token.New("test", token.IDENT, 0, 0),
					Value: "test",
				},
				Value: &ast.StringLiteral{
					Token: token.New("test", token.STRING, 0, 0),
					Value: "test",
				},
			},
		},
		{
			Name:       "It should parse constant bool declarations",
			Expression: "const test = true",
			ExpectedNode: &ast.ConstStatement{
				Token: token.New(token.CONST, token.CONST, 0, 0),
				Name: &ast.Identifier{
					Token: token.New("test", token.IDENT, 0, 0),
					Value: "test",
				},
				Value: &ast.BooleanLiteral{
					Token: token.New(token.TRUE, token.TRUE, 0, 0),
					Value: true,
				},
			},
		},
		{
			Name:       "It should parse constant array declarations",
			Expression: `const test = [1, "test", 't']`,
			ExpectedNode: &ast.ConstStatement{
				Token: token.New(token.CONST, token.CONST, 0, 0),
				Name: &ast.Identifier{
					Token: token.New("test", token.IDENT, 0, 0),
					Value: "test",
				},
				Value: &ast.ArrayLiteral{
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
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
