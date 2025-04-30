package obj

import "fmt"

// an identifier in Golflang.
// When used, acts as a placeholder for later lazy evaluation.
type Ident string

// Bool implements Obj.
func (i Ident) Bool() bool {
	return true
}

// Equal implements Obj.
func (i Ident) Equal(o Obj) bool {
	switch o := o.(type) {
	case Ident:
		return i == o
	default:
		return false
	}
}

// Hash implements Obj.
func (i Ident) Hash() Hash {
	return Hash(fmt.Sprintf("%#v", i))
}

// Kind implements Obj.
func (i Ident) Kind() ObjKind {
	return ObjKindIdent
}

// Repr implements Obj.
func (i Ident) Repr() string {
	return string(i)
}

// String implements Obj.
func (i Ident) String() string {
	return string(i)
}
