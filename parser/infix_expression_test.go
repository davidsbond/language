package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/token"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
	"github.com/stretchr/testify/assert"
)

func TestParser_InfixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name               string
		Expression         string
		ExpectedExpression *ast.InfixExpression
	}{
		{
			Name:       "It should parse addition",
			Expression: "1 + 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.PLUS,
				Token:    token.New(token.PLUS, token.PLUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse subtraction",
			Expression: "1 - 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.MINUS,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse multiplication",
			Expression: "1 * 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.ASTERISK,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse division",
			Expression: "1 / 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.SLASH,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse modulo",
			Expression: "2 % 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.MOD,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("2", token.NUMBER, 0, 0),
					Value: 2,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse greater than comparison",
			Expression: "2 > 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.GT,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("2", token.NUMBER, 0, 0),
					Value: 2,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse less than comparison",
			Expression: "2 < 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.LT,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("2", token.NUMBER, 0, 0),
					Value: 2,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse comparison",
			Expression: "2 == 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.EQUALS,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("2", token.NUMBER, 0, 0),
					Value: 2,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse not-comparison",
			Expression: "2 != 1",
			ExpectedExpression: &ast.InfixExpression{
				Operator: token.NOTEQ,
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Left: &ast.NumberLiteral{
					Token: token.New("2", token.NUMBER, 0, 0),
					Value: 2,
				},
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
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

			exp, ok := stmt.Expression.(*ast.InfixExpression)

			assert.True(t, ok)
			assert.Equal(t, tc.ExpectedExpression.String(), exp.String())
		})
	}
}
