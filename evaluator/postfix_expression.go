package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluatePostfixExpression(node *ast.PostfixExpression, scope *object.Scope) object.Object {
	val := scope.Get(node.Left.Value)
	newVal := evaluatePostfix(node.Operator, val)

	if isError(newVal) {
		return newVal
	}

	switch obj := val.(type) {
	case *object.Constant:
		return object.Error("cannot modify constant '%s'", node.Left.Value)
	case *object.Atomic:
		obj.Set(newVal)
		return scope.Set(node.Left.Value, newVal)
	default:
		return scope.Set(node.Left.Value, newVal)
	}
}

func evaluatePostfix(operator string, obj object.Object) object.Object {
	num, err := getNumberFromObject(obj)

	if err != nil {
		return object.Error(err.Error())
	}

	switch operator {
	case token.INC:
		return &object.Number{Value: num.Value + 1}
	case token.DEC:
		return &object.Number{Value: num.Value - 1}
	default:
		return object.Error("type Number does not support operator %s", operator)
	}
}
