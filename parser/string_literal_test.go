package parser_test

import (
	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
)

func TestParser_StringLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name string
		Expression        string
		ExpectedLiteral *ast.StringLiteral
	}{
		{
			Name: "It should parse string literals",
			Expression:        `"test"`,
			ExpectedLiteral: &ast.StringLiteral {
				Token: token.New("1", token.STRING, 0, 0),
				Value: "test",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result := parser.Parse()

			assert.Len(t, result.Nodes, 1)
			
			stmt, ok := result.Nodes[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			lit, ok := stmt.Expression.(*ast.StringLiteral)
			assert.True(t, ok)
			
			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
