package parser

import "github.com/davidsbond/dave/ast"

func (p *Parser) parseIdentifier() ast.Node {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}
