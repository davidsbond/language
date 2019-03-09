package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The AsyncStatement type represents an async function declaration in
	// therce code.
	//
	// Example:
	// async func add(a, b) {
	// return a + b
	//}
	AsyncStatement struct {
		Token *token.Token
		Value Node
	}
)

func (as *AsyncStatement) String() string {
	return fmt.Sprintf("async %s", as.Value.String())
}
