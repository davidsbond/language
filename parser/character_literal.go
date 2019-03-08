package parser

import (
	"unicode/utf8"

	"github.com/davidsbond/dave/ast"
)

func (p *Parser) parseCharacterLiteral() ast.Node {
	r, _ := utf8.DecodeRuneInString(p.currentToken.Literal)

	return &ast.CharacterLiteral{
		Token: p.currentToken,
		Value: r,
	}
}
