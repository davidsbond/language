package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseCallExpression(function ast.Node) ast.Node {
	exp := &ast.CallExpression{
		Token:    p.currentToken,
		Function: function,
	}

	exp.Arguments = p.parseExpressionList(token.RPAREN)

	return exp
}
