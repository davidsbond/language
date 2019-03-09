package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseAsyncStatement() ast.Node {
	stmt := &ast.AsyncStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.FUNCTION) {
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
