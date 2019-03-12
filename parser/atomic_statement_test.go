package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
)

func TestParser_AtomicStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name              string
		Expression        string
		ExpectedStatement *ast.AtomicStatement
	}{
		{
			Name:       "It should parse atomic number declarations",
			Expression: "atomic test = 1",
			ExpectedStatement: &ast.AtomicStatement{
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
			Name:       "It should parse atomic string declarations",
			Expression: `atomic test = "test"`,
			ExpectedStatement: &ast.AtomicStatement{
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
			Name:       "It should parse atomic bool declarations",
			Expression: "atomic test = true",
			ExpectedStatement: &ast.AtomicStatement{
				Token: token.New(token.ATOMIC, token.ATOMIC, 0, 0),
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
			Name:       "It should parse atomic array declarations",
			Expression: `atomic test = [1, "test", 't']`,
			ExpectedStatement: &ast.AtomicStatement{
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
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result, _ := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			stmt, ok := result.Nodes[0].(*ast.AtomicStatement)

			assert.True(t, ok)
			assert.Equal(t, tc.ExpectedStatement.String(), stmt.String())
		})
	}
}
