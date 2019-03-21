package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_InfixExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse addition",
			Expression: "1 + 1",
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
			ExpectedNode: &ast.InfixExpression{
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
		{
			Name:       "It should parse less than or equal to",
			Expression: "2 <= 1",
			ExpectedNode: &ast.InfixExpression{
				Operator: token.LTEQ,
				Token:    token.New(token.LTEQ, token.LTEQ, 0, 0),
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
			Name:       "It should parse greater than or equal to",
			Expression: "2 >= 1",
			ExpectedNode: &ast.InfixExpression{
				Operator: token.GTEQ,
				Token:    token.New(token.GTEQ, token.GTEQ, 0, 0),
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
		tc.Run(t)
	}
}
