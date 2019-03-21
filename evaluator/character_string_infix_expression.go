package evaluator

import (
	"github.com/davidsbond/language/object"
	"github.com/davidsbond/language/token"
)

func characterStringInfixExpression(operator string, left, right object.Object) object.Object {
	var l string
	var r string

	switch obj := left.(type) {
	case *object.String:
		l = obj.Value
	case *object.Character:
		l = obj.String()
	case *object.Constant:
		return characterStringInfixExpression(operator, obj.Value, right)
	case *object.Atomic:
		return characterStringInfixExpression(operator, obj.Value(), right)
	default:
		return object.Error("type %s does not support operator %s", obj.Type(), operator)
	}

	switch obj := right.(type) {
	case *object.String:
		r = obj.Value
	case *object.Character:
		r = obj.String()
	case *object.Constant:
		return characterStringInfixExpression(operator, left, obj.Value)
	case *object.Atomic:
		return characterStringInfixExpression(operator, left, obj.Value())
	default:
		return object.Error("type %s does not support operator %s", obj.Type(), operator)
	}

	switch operator {
	case token.PLUS:
		return &object.String{Value: l + r}
	default:
		return object.Error("types %s and %s do not support operator %s", left.Type(), right.Type(), operator)
	}
}
