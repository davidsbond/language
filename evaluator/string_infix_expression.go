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
		return object.Error("type String does not support operator %s", operator)
	case token.PLUS:
		return &object.String{Value: trueLeft.Value + trueRight.Value}
	case token.EQUALS:
		if trueLeft.Value == trueRight.Value {
			return TRUE
		}

		return FALSE
	case token.NOTEQ:
		if trueLeft.Value != trueRight.Value {
			return TRUE
		}

		return FALSE
	}
}

func getStringFromObject(obj object.Object) (*object.String, error) {
	switch val := obj.(type) {
	case *object.Constant:
		return getStringFromObject(val.Value)
	case *object.Atomic:
		return getStringFromObject(val.Value())
	case *object.String:
		return obj.(*object.String), nil
	case nil:
		return nil, fmt.Errorf("cannot cast nil value")
	default:
		return nil, fmt.Errorf("cannot cast type %s to type String", obj.Type())
	}
}
