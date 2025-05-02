//go:generate enumer -type=ObjKind
package obj

import "fmt"

// ObjKind helps determine what kind of object it is
type ObjKind uint

const (
	ObjKindNone ObjKind = iota
	ObjKindInt
	ObjKindDec
	ObjKindBool
	ObjKindStr
	ObjKindList
	ObjKindMap
	ObjKindIdent
	ObjKindLambda
)

// Obj is an interface for all values in golflang.
// Any values must implement this interface.
type Obj interface {
	Bool() bool
	Equal(o Obj) bool
	Kind() ObjKind  // the kind of this object
	Repr() string   // debug representation of this object
	String() string // stringified representation of this object, for printing
	Hash() Hash     // fmt.Sprintf("%#v") for implementing Map
}

type Hash string // for implementing Map

/* === force implementation of Object === */

var _ Obj = none
var _ Obj = ZeroInt()
var _ Obj = ZeroDec()
var _ Obj = ZeroBool()
var _ Obj = ZeroStr()
var _ Obj = ZeroList()
var _ Obj = ZeroMap()
var _ Obj = Ident("")

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
