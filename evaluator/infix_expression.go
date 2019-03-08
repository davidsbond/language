package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateInfixExpression(node *ast.InfixExpression, scope *object.Scope) object.Object {
	left := Evaluate(node.Left, scope)

	if left.Type() == object.TypeError {
		return left
	}

	right := Evaluate(node.Right, scope)

	if right.Type() == object.TypeError {
		return right
	}

	switch {
	case left.Type() == object.TypeNumber && right.Type() == object.TypeNumber:
		return evaluateNumberInfixExpression(node.Operator, left, right)
	case left.Type() == object.TypeString && right.Type() == object.TypeString:
		return evaluateStringInfixExpression(node.Operator, left, right)
	case left.Type() == object.TypeCharacter && right.Type() == object.TypeCharacter:
		return evaluateCharacterInfixExpression(node.Operator, left, right)
	default:
		return object.Error("types %s and %s do not support operator %s", left.Type(), right.Type(), node.Operator)
	}
}
