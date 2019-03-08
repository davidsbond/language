package object

import (
	"fmt"
)

const (
	// TypeError is the type for an error value.
	TypeError = "error"
)

type (
	// The Error type represents a parser error
	err struct {
		message string
		line    int
		column  int
	}
)

// Error creates a new error in memory using the given token to note the line/column.
// Standard library style message formatting can be used to create a formatted error
// message.
func Error(msg string, args ...interface{}) Object {
	return &err{
		message: fmt.Sprintf(msg, args...),
	}
}

func (e *err) Type() Type {
	return TypeError
}

func (e *err) Clone() Object {
	return &err{
		message: e.message,
		line:    e.line,
		column:  e.column,
	}
}

func (e *err) String() string {
	return e.message
}
