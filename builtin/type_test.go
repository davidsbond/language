package builtin_test

import (
	"testing"

	"github.com/davidsbond/dave/builtin"
	"github.com/davidsbond/dave/object"
)

func TestBuiltin_Type(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Args           []object.Object
		ExpectedObject object.Object
	}{
		{
			Name:           "It should return the type for a valid argument",
			Args:           []object.Object{&object.Number{}},
			ExpectedObject: &object.String{Value: "Number"},
		},
		{
			Name:           "It should return an error invalid arguments",
			Args:           []object.Object{&object.Number{}, &object.Number{}},
			ExpectedObject: object.Error("built-in 'type' function only takes one argument"),
		},
		{
			Name:           "It should return an error invalid arguments",
			Args:           []object.Object{},
			ExpectedObject: object.Error("built-in 'type' function only takes one argument"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := builtin.Type(tc.Args...)

			assertEqualObjects(t, tc.ExpectedObject, result)
		})
	}
}
