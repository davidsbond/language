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

func TestParser_HashLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.HashLiteral
	}{
		{
			Name: "It should parse valid hash literals",
			Expression: `{ 
				"a": 1, 
				"b": "test", 
				"c": 't' 
			}`,
			ExpectedLiteral: &ast.HashLiteral{
				Token: token.New(token.LBRACE, token.LBRACE, 0, 0),
				Pairs: map[ast.Node]ast.Node{},
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

			stmt, ok := result.Nodes[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			lit, ok := stmt.Expression.(*ast.HashLiteral)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
