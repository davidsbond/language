package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseFunctionLiteral() ast.Node {
	lit := &ast.FunctionLiteral{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	lit.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	results := []*ast.Identifier{}

	// If there are no parameters, return the empty array and move
	// on
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return results
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	ident := &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	results = append(results, ident)

	for p.peekTokenIs(token.COMMA) {
		// We move twice here so that currentToken is always the
		// identifier and peekToken is always either the comma
		// or closing parenthesis
		p.nextToken()

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		ident := &ast.Identifier{
			Token: p.currentToken,
			Value: p.currentToken.Literal,
		}

		results = append(results, ident)

	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return results
}
