package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
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
