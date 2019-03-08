package evaluator

import (
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluateCharacterInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.Character instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft := getCharacterFromObject(left)
	trueRight := getCharacterFromObject(right)

	switch operator {
	default:
		return object.Error("type %s does not support operator %s", trueLeft.Type(), operator)
	case token.PLUS:
		return &object.String{Value: string([]rune{trueLeft.Value, trueRight.Value})}
	case token.MINUS:
		return &object.Character{Value: trueLeft.Value - trueRight.Value}
	case token.ASTERISK:
		return &object.Character{Value: trueLeft.Value * trueRight.Value}
	case token.SLASH:
		return &object.Character{Value: trueLeft.Value / trueRight.Value}
	case token.LT:
		return &object.Boolean{Value: trueLeft.Value < trueRight.Value}
	case token.GT:
		return &object.Boolean{Value: trueLeft.Value > trueRight.Value}
	case token.EQUALS:
		return &object.Boolean{Value: trueLeft.Value == trueRight.Value}
	}
}

func getCharacterFromObject(obj object.Object) *object.Character {
	switch val := obj.(type) {
	case *object.Constant:
		return val.Value.(*object.Character)
	case *object.Atomic:
		return val.Value().(*object.Character)
	default:
		return obj.(*object.Character)
	}
}
