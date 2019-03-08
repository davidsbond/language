package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
)

func TestParser_VarStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name              string
		Expression        string
		ExpectedStatement *ast.VarStatement
	}{
		{
			Name:       "It should parse number variable declarations",
			Expression: "var test = 1",
			ExpectedStatement: &ast.VarStatement{
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
			ExpectedStatement: &ast.VarStatement{
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
			ExpectedStatement: &ast.VarStatement{
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
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result, _ := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			stmt, ok := result.Nodes[0].(*ast.VarStatement)

			assert.True(t, ok)
			assert.Equal(t, tc.ExpectedStatement.String(), stmt.String())
		})
	}
}
