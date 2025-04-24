package obj

// ObjKind helps determine what kind of object it is
type ObjKind uint

const (
	KindNone ObjKind = iota
	KindInt
	KindFloat
	KindBool
	KindString
	KindList
	KindMap
	KindBuiltinFunc
	KindAlias
)

// Obj is an interface for all values in golflang.
// Any values musst implement this interface.
type Obj interface {
	Kind() ObjKind // the kind of this object
	String() string
	Bool() bool
	Equal(o Obj) bool
}

/* === force implementation of Object === */

var _ Obj = Int(0)
var _ Obj = Float(0.0)
var _ Obj = Bool(false)
var _ Obj = String("")

// var _ Obj = List([]Obj{})
// var _ Obj = Map(make(map[Obj]Obj))
// var _ Obj = BuiltinFunc()
// var _ Obj = Alias()
