// Package token contains constants used to identify collections of
// lexemes within the source code.
package token

const (
	// IDENT is the token type used for an identifier
	IDENT = "identifier"

	// NUMBER is the token type used for a number
	NUMBER = "number"

	// EOF is the token type used when the end of the source
	// code is reached.
	EOF = "EOF"

	// VAR is the token type used when declaring a mutable
	// variable
	VAR = "var"

	// CONST is the token type used when declaring an immutable
	// variable.
	CONST = "const"

	// ATOMIC is the token type used when declaring an atomic variable
	ATOMIC = "atomic"

	// ASSIGN is the token type used for assignment statements
	ASSIGN = "="

	// EQUALS is the token type used for equality statements.
	EQUALS = "=="

	// NOTEQ is the token type used for not equals operators
	NOTEQ = "!="

	// BANG is the token type used for not expressions.
	BANG = "!"

	// STRING is the token type used for opening/closing literal strings.
	STRING = `"`

	// CHAR is the token type used for opening/closing literal characters.
	CHAR = `'`

	// PLUS is the token type used for addition
	PLUS = "+"

	// MINUS is the token type used for subtraction/negative numbers
	MINUS = "-"

	// LT is the token type used for the less than symbol
	LT = "<"

	// GT is the token type used for the greater than symbol
	GT = ">"

	// GTEQ is the token type used for the greater than or equal to symbol
	GTEQ = ">="

	// LTEQ is the token type used for the less than or equal to symbol
	LTEQ = "<="

	// ASTERISK is the token type used for the '*' symbol.
	ASTERISK = "*"

	// SLASH is the token type used for the '/' symbol.
	SLASH = "/"

	// MOD is the token type used for the percent symbol.
	MOD = "%"

	// TRUE is the token type used for a boolean 'true'
	TRUE = "true"

	// FALSE is the token type used for a boolean 'false'
	FALSE = "false"

	// RETURN is the token type used for a return statement
	RETURN = "return"

	// FUNCTION is the token type used for a function statement
	FUNCTION = "func"

	// LPAREN is the token type used for a left-parenthesis symbol
	LPAREN = "("

	// RPAREN is the token type used for a right-parenthesis symbol
	RPAREN = ")"

	// COMMA is the token type used for ',' characters.
	COMMA = ","

	// LBRACE is the token type used for '{' characters
	LBRACE = "{"

	// RBRACE is the token type used for '}' characters
	RBRACE = "}"

	// ASYNC is the token type used for an async statement
	ASYNC = "async"

	// AWAIT is the token type used for an await statement
	AWAIT = "await"

	// COLON is the token type used for colon characters
	COLON = ":"

	// LBRACKET is the token type used for a left bracket
	LBRACKET = "["

	// RBRACKET is the token type used for a right bracket
	RBRACKET = "]"

	// COMMENT is the token type for a single-line comment.
	COMMENT = "//"

	// POW is the token type used for a power-of character.
	POW = "^"

	// INC is the token type used for incrementing
	INC = "++"

	// DEC is the token type used for decrementing
	DEC = "--"

	// SQRT is the token type for square-root characters.
	SQRT = "√"

	// IF is the token type for if statements
	IF = "if"

	// ELSE is the token type for else statements
	ELSE = "else"
)

var (
	keywords = map[string]Type{
		"async":  ASYNC,
		"atomic": ATOMIC,
		"await":  AWAIT,
		"const":  CONST,
		"func":   FUNCTION,
		"false":  FALSE,
		"return": RETURN,
		"true":   TRUE,
		"var":    VAR,
		"if":     IF,
		"else":   ELSE,
	}
)

type (
	// The Type type is used to differentiate between tokens
	Type string

	// The Token type represents a collection of lexemes
	// within the source code.
	Token struct {
		Type    Type
		Literal string
		Line    int
		Column  int
	}
)

// New creates a new instance of the Token type using the provided literal, token
// type, line & column numbers.
func New(literal string, t Type, l, c int) *Token {
	return &Token{
		Literal: literal,
		Type:    t,
		Line:    l,
		Column:  c,
	}
}

func (t *Token) String() string {
	return t.Literal
}

// LookupIdentifier checks to see if a given identifier has a matching
// token type. These are usually used for keywords.
func LookupIdentifier(ident string) (t Type) {
	t, ok := keywords[ident]

	if !ok {
		t = IDENT
	}

	return
}
