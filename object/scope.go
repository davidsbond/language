package object

type (
	// The Scope type contains all accessible objects for a certain scope. It allows for
	// objects with the same name to be declared at different scopes and access them in order
	// of current scope to last parent scope.
	Scope struct {
		objects map[string]Object
		parent  *Scope
	}
)

// NewScope creates a new instance of the Scope type.
func NewScope() *Scope {
	return &Scope{
		objects: make(map[string]Object),
	}
}

// NewChildScope creates a new scope using the called scope as the
// parent.
func (s *Scope) NewChildScope() *Scope {
	scp := NewScope()
	scp.parent = s

	return scp
}

// Set attempts to set a value in the current scope using the given
// name.
func (s *Scope) Set(name string, val Object) Object {
	obj := s.Get(name)

	if obj == nil {
		s.objects[name] = val
		return val
	}

	switch object := obj.(type) {
	case *Constant:
		panic("can't change a constant")
	case *Atomic:
		object.Set(val)
	}

	return val
}

// Get attempts to obtain an object from the scope using the given name. If it
// does not exist at the current scope, the parent is checked until the value
// is found.
func (s *Scope) Get(name string) Object {
	obj, ok := s.objects[name]

	if !ok && s.parent != nil {
		obj = s.parent.Get(name)
	}

	return obj
}
