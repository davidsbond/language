package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseAwaitStatement() ast.Node {
	stmt := &ast.AwaitStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
