package evaluator_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/language/evaluator"
	"github.com/davidsbond/language/lexer"
	"github.com/davidsbond/language/object"
	"github.com/davidsbond/language/parser"
)

type (
	EvaluatorTest struct {
		Name           string
		Expression     string
		ExpectedObject object.Object
	}
)

func TestEvaluator_Evaluate(t *testing.T) {
	t.Parallel()

	tt := []EvaluatorTest{
		{
			Name: "It should evaluate fibonnaci",
			ExpectedObject: &object.Number{
				Value: 34,
			},
			Expression: `
			func fib(n) {
				if (n == 1) {
					return n
				}

				if (n < 1) {
					return n
				}

				return fib(n - 1) + fib(n - 2)
			}

			fib(9)
			`,
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}

func (et *EvaluatorTest) Run(t *testing.T) {
	t.Run(et.Name, func(t *testing.T) {
		rd := bufio.NewReader(strings.NewReader(et.Expression))
		lex, _ := lexer.New(rd)
		parser := parser.New(lex)
		ast, _ := parser.Parse()

		scope := object.NewScope()
		actual := evaluator.Evaluate(ast, scope)

		switch expected := et.ExpectedObject.(type) {
		default:
			assertEqualObjects(t, expected, actual)
		case *object.Hash:
			assertEqualHashes(t, expected, actual)
		}
	})
}

func assertEqualHashes(t *testing.T, expected *object.Hash, actual object.Object) {
	byKey := make(map[float64]object.Object)
	for key, val := range actual.(*object.Hash).Pairs {
		byKey[key.Value] = val.Value
	}

	for key, expected := range expected.Pairs {
		actual, _ := byKey[key.Value]

		if expected.Value.Type() != actual.Type() {
			t.Fatalf("expected object type %s, got %s", expected.Value.Type(), actual.Type())
		}

		if expected.Value.String() != actual.String() {
			t.Fatalf("expected %s, got %s", expected.Value.String(), actual.String())
		}
	}
}

func assertEqualObjects(t *testing.T, expected, actual object.Object) {
	if expected.Type() != actual.Type() {
		t.Fatalf("expected object type %s, got %s", expected.Type(), actual.Type())
	}

	if expected.String() != actual.String() {
		t.Fatalf("expected %s, got %s", expected.String(), actual.String())
	}
}
