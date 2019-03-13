package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func functionLiteral(node *ast.FunctionLiteral, scope *object.Scope) object.Object {
	lit := &object.Function{
		Name:       node.Name,
		Parameters: node.Parameters,
		Body:       node.Body,
	}

	scope.Set(lit.Name.Value, lit)

	return lit
}
