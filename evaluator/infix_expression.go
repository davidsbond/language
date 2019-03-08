package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateInfixExpression(node *ast.InfixExpression, scope *object.Scope) object.Object {
	left := Evaluate(node.Left, scope)

	// check err

	right := Evaluate(node.Right, scope)

	// check err

	switch {
	case left.Type() == object.TypeNumber && right.Type() == object.TypeNumber:
		return evaluateNumberInfixExpression(node.Operator, left, right)
	case left.Type() == object.TypeString && right.Type() == object.TypeString:
		return evaluateStringInfixExpression(node.Operator, left, right)
	case left.Type() == object.TypeCharacter && right.Type() == object.TypeCharacter:
		return evaluateCharacterInfixExpression(node.Operator, left, right)
	}

	return nil
}
