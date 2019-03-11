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

			byString := make(map[string]ast.Node)
			for key, val := range lit.Pairs {
				byString[key.String()] = val
			}

			for key, expected := range tc.ExpectedLiteral.Pairs {
				actual, ok := byString[key.String()]

				assert.True(t, ok)
				assert.Equal(t, expected.String(), actual.String())
			}
		})
	}
}
