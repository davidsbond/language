package builtin

import (
	"github.com/davidsbond/language/object"
)

// Type is a built-in function that returns a string representation of an
// object's type. Returns an error if too many arguments are provided.
func Type(args ...object.Object) object.Object {
	if len(args) > 1 || len(args) == 0 {
		return object.Error("built-in 'type' function only takes one argument")
	}

	return &object.String{
		Value: string(args[0].Type()),
	}
}
