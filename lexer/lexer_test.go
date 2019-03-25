package lexer_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/language/lexer"
	"github.com/davidsbond/language/token"
)

func TestLexer_NextToken(t *testing.T) {
	tt := []struct {
		Expression    string
		ExpectedToken *token.Token
	}{
		{
			Expression: "!=",
			ExpectedToken: &token.Token{
				Literal: token.NOTEQ,
				Type:    token.NOTEQ,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "// A comment",
			ExpectedToken: &token.Token{
				Literal: " A comment",
				Type:    token.COMMENT,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "^",
			ExpectedToken: &token.Token{
				Literal: token.POW,
				Type:    token.POW,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "∞",
			ExpectedToken: &token.Token{
				Literal: "∞",
				Type:    token.IDENT,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "√",
			ExpectedToken: &token.Token{
				Literal: token.SQRT,
				Type:    token.SQRT,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "!",
			ExpectedToken: &token.Token{
				Literal: token.BANG,
				Type:    token.BANG,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: ":",
			ExpectedToken: &token.Token{
				Literal: token.COLON,
				Type:    token.COLON,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "\n",
			ExpectedToken: &token.Token{
				Literal: token.EOF,
				Type:    token.EOF,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "const",
			ExpectedToken: &token.Token{
				Literal: "const",
				Type:    token.CONST,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "atomic",
			ExpectedToken: &token.Token{
				Literal: "atomic",
				Type:    token.ATOMIC,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "var",
			ExpectedToken: &token.Token{
				Literal: "var",
				Type:    token.VAR,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "await",
			ExpectedToken: &token.Token{
				Literal: "await",
				Type:    token.AWAIT,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "async",
			ExpectedToken: &token.Token{
				Literal: "async",
				Type:    token.ASYNC,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "true",
			ExpectedToken: &token.Token{
				Literal: "true",
				Type:    token.TRUE,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "false",
			ExpectedToken: &token.Token{
				Literal: "false",
				Type:    token.FALSE,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "=",
			ExpectedToken: &token.Token{
				Literal: "=",
				Type:    token.ASSIGN,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "==",
			ExpectedToken: &token.Token{
				Literal: "==",
				Type:    token.EQUALS,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "identifier",
			ExpectedToken: &token.Token{
				Literal: "identifier",
				Type:    token.IDENT,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: "123.456",
			ExpectedToken: &token.Token{
				Literal: "123.456",
				Type:    token.NUMBER,
				Line:    1,
				Column:  0,
			},
		},
		{
			Expression: `"test"`,
			ExpectedToken: &token.Token{
				Literal: "test",
				Type:    token.STRING,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: `'a'`,
			ExpectedToken: &token.Token{
				Literal: "a",
				Type:    token.CHAR,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "+",
			ExpectedToken: &token.Token{
				Literal: "+",
				Type:    token.PLUS,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "-",
			ExpectedToken: &token.Token{
				Literal: "-",
				Type:    token.MINUS,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "*",
			ExpectedToken: &token.Token{
				Literal: "*",
				Type:    token.ASTERISK,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "/",
			ExpectedToken: &token.Token{
				Literal: "/",
				Type:    token.SLASH,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "%",
			ExpectedToken: &token.Token{
				Literal: "%",
				Type:    token.MOD,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "<",
			ExpectedToken: &token.Token{
				Literal: "<",
				Type:    token.LT,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: ">",
			ExpectedToken: &token.Token{
				Literal: ">",
				Type:    token.GT,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "{",
			ExpectedToken: &token.Token{
				Literal: "{",
				Type:    token.LBRACE,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "}",
			ExpectedToken: &token.Token{
				Literal: "}",
				Type:    token.RBRACE,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "(",
			ExpectedToken: &token.Token{
				Literal: "(",
				Type:    token.LPAREN,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: ")",
			ExpectedToken: &token.Token{
				Literal: ")",
				Type:    token.RPAREN,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: ",",
			ExpectedToken: &token.Token{
				Literal: ",",
				Type:    token.COMMA,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "return",
			ExpectedToken: &token.Token{
				Literal: "return",
				Type:    token.RETURN,
				Line:    2,
				Column:  0,
			},
		},
		{
			Expression: "[",
			ExpectedToken: &token.Token{
				Literal: token.LBRACKET,
				Type:    token.LBRACKET,
				Line:    0,
				Column:  0,
			},
		},
		{
			Expression: "]",
			ExpectedToken: &token.Token{
				Literal: token.RBRACKET,
				Type:    token.RBRACKET,
				Line:    0,
				Column:  0,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Expression, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))

			lex, _ := lexer.New(rd)

			expected := tc.ExpectedToken
			actual, _ := lex.NextToken()

			if expected.Literal != actual.Literal {
				t.Fatalf("expected literal %s, got %s", expected.Literal, actual.Literal)
			}

			if expected.Type != actual.Type {
				t.Fatalf("expected type %s, got %s", expected.Type, actual.Type)
			}
		})
	}
}
