package builtin

import (
	"github.com/davidsbond/language/object"
)

// Len returns the length of a given string, array or hash.
func Len(args ...object.Object) object.Object {
	if len(args) > 1 || len(args) == 0 {
		return object.Error("built-in 'len' function only takes one argument")
	}

	switch obj := args[0].(type) {
	default:
		return object.Error("built-in 'len' does not support type %s", obj.Type())
	case *object.Array:
		return &object.Number{Value: float64(len(obj.Elements))}
	case *object.String:
		return &object.Number{Value: float64(len(obj.Value))}
	case *object.Constant:
		return Len(obj.Value)
	case *object.Atomic:
		return Len(obj.Value())
	}
}
