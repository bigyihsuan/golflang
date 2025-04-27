package obj

import "fmt"

// an identifier in Golflang.
// When used, acts as a placeholder for later lazy evaluation.
type Ident string

// Bool implements Obj.
// TODO: true if the identifier has already been declared/defined in this scope.
func (i Ident) Bool() bool {
	return false
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
	return fmt.Sprintf("ident(%s)", string(i))
}

// String implements Obj.
func (i Ident) String() string {
	return string(i)
}
