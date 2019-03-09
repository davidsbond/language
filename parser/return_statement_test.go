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

func TestParser_ReturnStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.ReturnStatement
	}{
		{
			Name:       "It should parse string literals",
			Expression: `return "test"`,
			ExpectedLiteral: &ast.ReturnStatement{
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
			ExpectedLiteral: &ast.ReturnStatement{
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
			ExpectedLiteral: &ast.ReturnStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				ReturnValue: &ast.CharacterLiteral{
					Token: token.New("'", token.CHAR, 0, 0),
					Value: '1',
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

			stmt, ok := result.Nodes[0].(*ast.ReturnStatement)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), stmt.String())
		})
	}
}
