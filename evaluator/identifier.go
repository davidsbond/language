package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func identifier(node *ast.Identifier, scope *object.Scope) object.Object {
	// Check if a special value exists with this name
	if value, ok := values[node.Value]; ok {
		return value
	}

	// Check for variables stored in the scope
	if val := scope.Get(node.Value); val != nil {
		return val
	}

	// Otherwise look for a built-in function
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return object.Error("cannot find value for identifier %s", node.Value)
}
