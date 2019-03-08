package evaluator

import (
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluateStringInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.String instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft := getStringFromObject(left)
	trueRight := getStringFromObject(right)

	switch operator {
	default:
		return nil
	case token.PLUS:
		return &object.String{Value: trueLeft.Value + trueRight.Value}
	case token.EQUALS:
		return &object.Boolean{Value: trueLeft.Value == trueRight.Value}
	}
}

func getStringFromObject(obj object.Object) *object.String {
	switch val := obj.(type) {
	case *object.Constant:
		return val.Value.(*object.String)
	case *object.Atomic:
		return val.Value().(*object.String)
	default:
		return obj.(*object.String)
	}
}
