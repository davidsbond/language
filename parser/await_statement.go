package parser

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
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
