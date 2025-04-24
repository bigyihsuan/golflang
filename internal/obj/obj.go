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
	AsBool() bool
	Equal(o Obj) bool
}

/* === force implementation of Object === */

var _ Obj = Int(0)
var _ Obj = Float(0.0)
var _ Obj = Bool(false)

// var _ Object = String("")
// var _ Object = List([]Object{})
// var _ Object = Map(make(map[Object]Object))
// var _ Object = BuiltinFunc()
// var _ Object = Alias()
