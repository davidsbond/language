package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateVarStatement(node *ast.VarStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	if isError(val) {
		return val
	}

	return scope.Set(node.Name.Value, val)
}
