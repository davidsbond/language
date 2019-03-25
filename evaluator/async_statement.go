package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func asyncStatement(node *ast.AsyncStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	if isError(val) {
		return val
	}

	switch fn := val.(type) {
	case *object.Function:
		fn.Async = true

		return scope.Set(fn.Name.Value, fn)
	default:
		return object.Error("expected type Function, got %s", fn.Type())
	}
}
