package obj

import "fmt"

type None struct{}

// Bool implements Obj.
func (n None) Bool() bool {
	return false
}

// Equal implements Obj.
func (n None) Equal(o Obj) bool {
	return false
}

// Hash implements Obj.
func (n None) Hash() Hash {
	return Hash(fmt.Sprintf("%#v", n))
}

// Kind implements Obj.
func (n None) Kind() ObjKind {
	return ObjKindNone
}

// Repr implements Obj.
func (n None) Repr() string {
	return "none"
}

// String implements Obj.
func (n None) String() string {
	return "none"
}

var none = None{}

func ZeroNone() None {
	return none
}
