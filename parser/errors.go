package parser

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The Error type represents a parser error
	Error struct {
		message string
		token   *token.Token
	}
)

func (p *Parser) error(msg string, args ...interface{}) {
	err := Error{
		message: fmt.Sprintf(msg, args...),
		token:   p.currentToken,
	}

	p.errors = append(p.errors, err)
}

func (err Error) Error() string {
	return fmt.Sprintf("(%d:%d): %s", err.token.Line, err.token.Column, err.message)
}
