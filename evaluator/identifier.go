package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateIdentifier(node *ast.Identifier, scope *object.Scope) object.Object {
	if val := scope.Get(node.Value); val != nil {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return object.Error("cannot find value for identifier %s", node.Value)
}
