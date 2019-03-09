package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/builtin"
	"github.com/davidsbond/dave/object"
)

var (
	functions = map[string]object.Builtin{
		"type": builtin.Type,
	}
)

func evaluateIdentifier(node *ast.Identifier, scope *object.Scope) object.Object {
	if val := scope.Get(node.Value); val != nil {
		return val
	}

	if builtin, ok := functions[node.Value]; ok {
		return builtin
	}

	return object.Error("cannot find value for identifier %s", node.Value)
}
