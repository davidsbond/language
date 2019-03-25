package ast

import (
	"fmt"

	"github.com/davidsbond/language/token"
)

type (
	// The AwaitStatement type represents an awaited call in
	// the source code
	//
	// Example:
	// await add(1, 2)
	AwaitStatement struct {
		Token *token.Token
		Value Node
	}
)

func (aw *AwaitStatement) String() string {
	return fmt.Sprintf("await %s", aw.Value.String())
}
