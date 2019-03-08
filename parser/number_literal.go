package parser

import (
	"strconv"

	"github.com/davidsbond/dave/ast"
)

func (p *Parser) parseNumberLiteral() ast.Node {
	lit := &ast.NumberLiteral{Token: p.currentToken}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)

	if err != nil {
		p.error(err.Error())
		return nil
	}

	lit.Value = value
	return lit
}
