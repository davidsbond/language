package builtin_test

import (
	"testing"

	"github.com/davidsbond/language/object"
)

func assertEqualObjects(t *testing.T, expected, actual object.Object) {
	if expected.Type() != actual.Type() {
		t.Fatalf("expected object type %s, got %s", expected.Type(), actual.Type())
	}

	if expected.String() != actual.String() {
		t.Fatalf("expected %s, got %s", expected.String(), actual.String())
	}
}
