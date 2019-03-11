package evaluator

import (
	"math"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateIndexExpression(node *ast.IndexExpression, scope *object.Scope) object.Object {
	left := Evaluate(node.Left, scope)

	if isError(left) {
		return left
	}

	index := Evaluate(node.Index, scope)

	if isError(index) {
		return index
	}

	container := getIndexableFromObject(left)

	if isError(container) {
		return container
	}

	switch obj := container.(type) {
	case *object.String:
		return evaluateStringIndexExpression(obj, index)
	case *object.Array:
		return evaluateArrayIndexExpression(obj, index)
	case *object.Hash:
		return evaluateHashIndexExpression(obj, index)
	default:
		return object.Error("type %s does not support indexing", container.Type())
	}
}

func evaluateHashIndexExpression(hash *object.Hash, index object.Object) object.Object {
	k, ok := index.(object.Hashable)

	if !ok {
		return object.Error("type %s cannot be used as a hash key", index.Type())
	}

	pair, ok := hash.Pairs[k.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}

func evaluateStringIndexExpression(str *object.String, index object.Object) object.Object {
	num, err := getNumberFromObject(index)

	if err != nil {
		return object.Error(err.Error())
	}

	idx := math.Floor(num.Value)
	if idx > float64(len(str.Value)) || idx < 0 {
		return object.Error("index out of range")
	}

	return &object.Character{Value: rune(str.Value[int(idx)])}
}

func evaluateArrayIndexExpression(arr *object.Array, index object.Object) object.Object {
	num, err := getNumberFromObject(index)

	if err != nil {
		return object.Error(err.Error())
	}

	idx := math.Floor(num.Value)
	if idx > float64(len(arr.Elements)) || idx < 0 {
		return object.Error("index out of range")
	}

	return arr.Elements[int(idx)]
}

func getIndexableFromObject(left object.Object) object.Object {
	switch obj := left.(type) {
	case *object.Constant:
		return getIndexableFromObject(obj.Value)
	case *object.Atomic:
		return getIndexableFromObject(obj.Value())
	case *object.Array:
		return obj
	case *object.Hash:
		return obj
	case *object.String:
		return obj
	default:
		return object.Error("type %s does not support indexing", obj.Type())
	}
}
