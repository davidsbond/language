package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	Comment struct {
		Token *token.Token
		Value string
	}
)

func (cm *Comment) String() string {
	return fmt.Sprintf("// %s", cm.Value)
}
