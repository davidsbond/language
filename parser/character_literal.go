package parser

import (
	"unicode/utf8"

	"github.com/davidsbond/language/ast"
)

func (p *Parser) parseCharacterLiteral() ast.Node {
	r, _ := utf8.DecodeRuneInString(p.currentToken.Literal)

	if r == utf8.RuneError {
		p.error("invalid character %s, characters must be utf-8", p.currentToken.Literal)
		return nil
	}

	return &ast.CharacterLiteral{
		Token: p.currentToken,
		Value: r,
	}
}
