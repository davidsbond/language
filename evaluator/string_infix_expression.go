package evaluator

import (
	"fmt"

	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluateStringInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.String instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft, err := getStringFromObject(left)

	if err != nil {
		return object.Error(err.Error())
	}

	trueRight, err := getStringFromObject(right)

	if err != nil {
		return object.Error(err.Error())
	}

	switch operator {
	default:
		return nil
	case token.PLUS:
		return &object.String{Value: trueLeft.Value + trueRight.Value}
	case token.EQUALS:
		return &object.Boolean{Value: trueLeft.Value == trueRight.Value}
	}
}

func getStringFromObject(obj object.Object) (*object.String, error) {
	switch val := obj.(type) {
	case *object.Constant:
		return val.Value.(*object.String), nil
	case *object.Atomic:
		return val.Value().(*object.String), nil
	case *object.String:
		return obj.(*object.String), nil
	case nil:
		return nil, fmt.Errorf("cannot cast nil value")
	default:
		return nil, fmt.Errorf("cannot cast type %s to type String", obj.Type())
	}
}
