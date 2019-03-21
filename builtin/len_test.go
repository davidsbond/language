package builtin_test

import (
	"testing"

	"github.com/davidsbond/language/builtin"
	"github.com/davidsbond/language/object"
)

func TestBuiltin_Len(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Args           []object.Object
		ExpectedObject object.Object
	}{
		{
			Name:           "It should return the length of an array",
			Args:           []object.Object{&object.Array{Elements: []object.Object{&object.Number{}}}},
			ExpectedObject: &object.Number{Value: 1},
		},
		{
			Name:           "It should return the length of a string",
			Args:           []object.Object{&object.String{Value: "test"}},
			ExpectedObject: &object.Number{Value: 4},
		},
		{
			Name:           "It should return the length of a constant string",
			Args:           []object.Object{&object.Constant{Value: &object.String{Value: "test"}}},
			ExpectedObject: &object.Number{Value: 4},
		},
		{
			Name:           "It should return the length of an atomic string",
			Args:           []object.Object{object.MakeAtomic(&object.String{Value: "test"})},
			ExpectedObject: &object.Number{Value: 4},
		},
		{
			Name:           "It should return an error for invalid arguments",
			Args:           []object.Object{},
			ExpectedObject: object.Error("built-in 'len' function only takes one argument"),
		},
		{
			Name:           "It should return an error for invalid types",
			Args:           []object.Object{&object.Character{Value: 'a'}},
			ExpectedObject: object.Error("built-in 'len' does not support type Character"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := builtin.Len(tc.Args...)

			assertEqualObjects(t, tc.ExpectedObject, result)
		})
	}
}
