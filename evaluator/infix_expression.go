package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func infixExpression(node *ast.InfixExpression, scope *object.Scope) object.Object {
	left := Evaluate(node.Left, scope)

	if left.Type() == object.TypeError {
		return left
	}

	right := Evaluate(node.Right, scope)

	if right.Type() == object.TypeError {
		return right
	}

	if ident, ok := node.Left.(*ast.Identifier); node.Operator == token.ASSIGN && ok {
		if obj := scope.Get(ident.Value); obj != nil {
			switch val := obj.(type) {
			case *object.Constant:
				return object.Error("cannot reassign constant '%s'", ident.Value)
			case *object.Atomic:
				val.Set(right)
				return right
			}
		}

		return scope.Set(ident.Value, right)
	}

	switch {
	// 1 + 1
	case left.Type() == object.TypeNumber && right.Type() == object.TypeNumber:
		return numberInfixExpression(node.Operator, left, right)
	// "a" + "a"
	case left.Type() == object.TypeString && right.Type() == object.TypeString:
		return stringInfixExpression(node.Operator, left, right)
	// "a" + 'a'
	case left.Type() == object.TypeString && right.Type() == object.TypeCharacter:
		return characterStringInfixExpression(node.Operator, left, right)
	// 'a' + "a"
	case left.Type() == object.TypeCharacter && right.Type() == object.TypeString:
		return characterStringInfixExpression(node.Operator, left, right)
	// true == true
	case left.Type() == object.TypeBoolean && right.Type() == object.TypeBoolean:
		return booleanInfixExpression(node.Operator, left, right)
	// 'a' + 'a'
	case left.Type() == object.TypeCharacter && right.Type() == object.TypeCharacter:
		return characterInfixExpression(node.Operator, left, right)
	default:
		return object.Error("types %s and %s do not support operator %s", left.Type(), right.Type(), node.Operator)
	}
}
