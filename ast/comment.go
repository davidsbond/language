package ast

import (
	"fmt"

	"github.com/davidsbond/language/token"
)

type (
	// The Comment type represents a comment in the source code.
	Comment struct {
		Token *token.Token
		Value string
	}
)

func (cm *Comment) String() string {
	return fmt.Sprintf("// %s", cm.Value)
}
