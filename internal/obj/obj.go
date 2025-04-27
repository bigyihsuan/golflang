//go:generate stringer -type=ObjKind
package obj

// ObjKind helps determine what kind of object it is
type ObjKind uint

const (
	KindNone ObjKind = iota
	KindInt
	KindDec
	KindBool
	KindStr
	KindList
	KindMap
	KindBuiltinFunc
	KindIdent
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

var _ Obj = ZeroInt()
var _ Obj = ZeroDec()
var _ Obj = ZeroBool()
var _ Obj = ZeroStr()
var _ Obj = ZeroList()
var _ Obj = ZeroMap()

// var _ Obj = BuiltinFunc()
// var _ Obj = Ident()
