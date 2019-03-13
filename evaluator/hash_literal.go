package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func hashLiteral(node *ast.HashLiteral, scope *object.Scope) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for k, v := range node.Pairs {
		key := Evaluate(k, scope)

		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)

		if !ok {
			return object.Error("type %s cannot be used as hash key", key.Type())
		}

		value := Evaluate(v, scope)

		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()

		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}
