// Package lexer contains the lexical analyzer for the language. It breaks down the input
// source code into a sequence of lexemes to produce tokens
package lexer

import (
	"bufio"
	"io"
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	// The Lexer type breaks down the input source code into a sequence of lexemes and produces
	// tokens.
	Lexer struct {
		input   *bufio.Reader
		current rune
		column  int
		line    int
	}
)

// New creates a new instance of the Lexer type that will attempt to convert the data in the
// provided *bufio.Reader instance to tokens. Creating a new lexer will cause the first rune
// to be read, which may return an error if it is an invalid character or EOF.
func New(input *bufio.Reader) (lex *Lexer, err error) {
	lex = &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}

	err = lex.readRune()

	return
}

// NextToken converts the next set of lexemes in the source code to a token. Returns an error
// for invalid characters or if reading the source code fails.
func (l *Lexer) NextToken() (tok *token.Token, err error) {
	l.skipWhitespace()

	switch l.current {
	case 0:
		tok = token.New(token.EOF, token.EOF, 0, 0)
	case '+':
		tok = token.New(token.PLUS, token.PLUS, l.line, l.column)
	case '-':
		tok = token.New(token.MINUS, token.MINUS, l.line, l.column)
	case '<':
		tok = token.New(token.LT, token.LT, l.line, l.column)
	case '>':
		tok = token.New(token.GT, token.GT, l.line, l.column)
	case '/':
		tok = token.New(token.SLASH, token.SLASH, l.line, l.column)
	case '*':
		tok = token.New(token.ASTERISK, token.ASTERISK, l.line, l.column)
	case '%':
		tok = token.New(token.MOD, token.MOD, l.line, l.column)
	case '(':
		tok = token.New(token.LPAREN, token.LPAREN, l.line, l.column)
	case ')':
		tok = token.New(token.RPAREN, token.RPAREN, l.line, l.column)
	case ',':
		tok = token.New(token.COMMA, token.COMMA, l.line, l.column)
	case '{':
		tok = token.New(token.LBRACE, token.LBRACE, l.line, l.column)
	case '}':
		tok = token.New(token.RBRACE, token.RBRACE, l.line, l.column)
	case '[':
		tok = token.New(token.LBRACKET, token.LBRACKET, l.line, l.column)
	case ']':
		tok = token.New(token.RBRACKET, token.RBRACKET, l.line, l.column)
	case '=':
		var next rune

		next, err = l.peekRune()

		// Check if the next rune is an '=', for equality
		// operators.
		if next == '=' {
			err = l.readRune()

			// Create a new equals token
			tok = token.New(
				"==",
				token.EQUALS,
				l.line,
				l.column)
		} else {
			// Create a new assignment token
			tok = token.New(
				string(l.current),
				token.ASSIGN,
				l.line,
				l.column)
		}

	case '"':
		var runes []rune

		err = l.readRune()

		for l.current != '"' {
			runes = append(runes, l.current)
			err = l.readRune()
		}

		tok = token.New(
			string(runes),
			token.STRING,
			l.line,
			l.column,
		)

	case '\'':
		var runes []rune

		err = l.readRune()

		for l.current != '\'' {
			runes = append(runes, l.current)
			err = l.readRune()
		}

		tok = token.New(
			string(runes[0]),
			token.CHAR,
			l.line,
			l.column,
		)

	default:
		if isLetter(l.current) {
			var ident string

			// Read the identifier and check if it's a keyword.
			ident, err = l.readIdentifier()
			typ := token.LookupIdentifier(ident)
			tok = token.New(ident, typ, l.line, l.column)

		} else if isDigit(l.current) {
			var num string

			// Read the number and produce the token
			num, err = l.readNumber()
			tok = token.New(num, token.NUMBER, l.line, l.column)
		} else {
			err = l.error("unsupported character: %v", l.current)
		}
	}

	err = l.readRune()

	return
}

func (l *Lexer) readRune() (err error) {
	ch, _, err := l.input.ReadRune()

	if err == io.EOF {
		ch = 0
	}

	if ch == '\n' {
		l.column = 1
		l.line++
	}

	l.column++
	l.current = ch

	return
}

func (l *Lexer) peekRune() (ch rune, err error) {
	ch, _, err = l.input.ReadRune()

	if err != nil {
		return
	}

	err = l.input.UnreadRune()
	return
}

func (l *Lexer) readIdentifier() (ident string, err error) {
	var out strings.Builder

	for err == nil {
		var next rune

		out.WriteRune(l.current)

		next, err = l.peekRune()

		if isLetter(next) {
			err = l.readRune()
		} else {
			break
		}
	}

	ident = out.String()
	return
}

func (l *Lexer) readNumber() (num string, err error) {
	var out strings.Builder

	for err == nil {
		var next rune

		out.WriteRune(l.current)

		next, err = l.peekRune()

		if isDigit(next) {
			err = l.readRune()
		} else {
			break
		}
	}

	num = out.String()
	return
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func (l *Lexer) skipWhitespace() (err error) {
	for (l.current == ' ' ||
		l.current == '\t' ||
		l.current == '\n' ||
		l.current == '\r') &&
		err == nil {

		err = l.readRune()
	}

	return
}
