package builtin_test

import (
	"os"
	"testing"

	"github.com/davidsbond/language/builtin"
	"github.com/davidsbond/language/object"
)

func TestBuiltin_SetEnv(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name           string
		Args           []object.Object
		ExpectedObject object.Object
	}{
		{
			Name:           "It should set an env var",
			Args:           []object.Object{&object.String{Value: "TEST_ENV_VAR"}, &object.String{Value: "VALUE"}},
			ExpectedObject: &object.Null{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := builtin.SetEnv(tc.Args...)

			assertEqualObjects(t, tc.ExpectedObject, result)
		})
	}
}

func TestBuiltin_GetEnv(t *testing.T) {
	t.Parallel()
	os.Setenv("TEST_ENV_VAR", "VALUE")

	tt := []struct {
		Name           string
		Args           []object.Object
		ExpectedObject object.Object
	}{
		{
			Name:           "It should return the value of a set env var",
			Args:           []object.Object{&object.String{Value: "TEST_ENV_VAR"}},
			ExpectedObject: &object.String{Value: "VALUE"},
		},
		{
			Name:           "It should return blank for a non-set env var",
			Args:           []object.Object{&object.String{Value: "AOJFASOH"}},
			ExpectedObject: &object.String{Value: ""},
		},
		{
			Name:           "It should use constant strings",
			Args:           []object.Object{&object.Constant{Value: &object.String{Value: "TEST_ENV_VAR"}}},
			ExpectedObject: &object.String{Value: "VALUE"},
		},
		{
			Name:           "It should use atomic strings",
			Args:           []object.Object{object.MakeAtomic(&object.String{Value: "TEST_ENV_VAR"})},
			ExpectedObject: &object.String{Value: "VALUE"},
		},
		{
			Name:           "It should return an error for invalid arguments",
			Args:           []object.Object{},
			ExpectedObject: object.Error("built-in 'get_env' function only takes one argument"),
		},
		{
			Name:           "It should return an error for invalid argument types",
			Args:           []object.Object{&object.Number{Value: 1}},
			ExpectedObject: object.Error("built-in 'get_env' does not support type Number"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := builtin.GetEnv(tc.Args...)

			assertEqualObjects(t, tc.ExpectedObject, result)
		})
	}
}
