package object

import (
	"sync"
)

type (
	// The Atomic type is a wrapper around Object that indicates the
	// object is atomic. The value must be accessed via the getters/setters
	// that use a mutex.
	Atomic struct {
		value Object
		mutex sync.Mutex
	}
)

// MakeAtomic converts a given object into an atomic one
func MakeAtomic(obj Object) *Atomic {
	return &Atomic{
		value: obj,
		mutex: sync.Mutex{},
	}
}

// Type returns the type of the underlying object.
func (at *Atomic) Type() Type {
	at.mutex.Lock()
	defer at.mutex.Unlock()

	return at.value.Type()
}

// Set sets the atomic value
func (at *Atomic) Set(obj Object) {
	at.mutex.Lock()
	defer at.mutex.Unlock()

	at.value = obj
}

// Value returns a copy of the atomic value.
func (at *Atomic) Value() Object {
	at.mutex.Lock()
	defer at.mutex.Unlock()

	return at.value.Clone()
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (at *Atomic) Clone() Object {
	return at.value.Clone()
}

func (at *Atomic) String() string {
	at.mutex.Lock()
	defer at.mutex.Unlock()

	return at.value.String()
}
