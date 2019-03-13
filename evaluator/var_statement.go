package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func varStatement(node *ast.VarStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	if isError(val) {
		return val
	}

	switch val.(type) {
	case object.Builtin:
		return object.Error("built-in functions cannot be used as variables")
	default:
		return scope.Set(node.Name.Value, val)
	}
}
