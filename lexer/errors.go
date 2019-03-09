package lexer

import (
	"fmt"
)

type (
	// The Error type represents a parser error
	Error struct {
		message string
		line    int
		column  int
	}
)

func (l *Lexer) error(msg string, args ...interface{}) error {
	return Error{
		message: fmt.Sprintf(msg, args...),
		line:    l.line,
		column:  l.column,
	}
}

func (err Error) Error() string {
	return fmt.Sprintf("(%d:%d): %s", err.line, err.column, err.message)
}
