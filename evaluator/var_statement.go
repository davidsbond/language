package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func varStatement(node *ast.VarStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	if isError(val) {
		return val
	}

	return scope.Set(node.Name.Value, val)
}
