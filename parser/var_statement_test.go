package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_VarStatement(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse number variable declarations",
			Expression: "var test = 1",
			ExpectedNode: &ast.VarStatement{
				Token: token.New(token.VAR, token.VAR, 0, 0),
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
			Name:       "It should parse string variable declarations",
			Expression: `var test = "test"`,
			ExpectedNode: &ast.VarStatement{
				Token: token.New(token.VAR, token.VAR, 0, 0),
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
			Name:       "It should parse variable bool declarations",
			Expression: "var test = true",
			ExpectedNode: &ast.VarStatement{
				Token: token.New(token.VAR, token.VAR, 0, 0),
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
			Name:       "It should parse variable array declarations",
			Expression: `var test = [1, "test", 't']`,
			ExpectedNode: &ast.VarStatement{
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
