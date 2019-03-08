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

func TestParser_CharacterLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.CharacterLiteral
	}{
		{
			Name:       "It should parse character literals",
			Expression: `'t'`,
			ExpectedLiteral: &ast.CharacterLiteral{
				Token: token.New("t", token.CHAR, 0, 0),
				Value: 't',
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

			lit, ok := stmt.Expression.(*ast.CharacterLiteral)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
