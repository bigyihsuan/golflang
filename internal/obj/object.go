package obj

// ObjectKind helps determine what kind of object it is
type ObjectKind uint

const (
	KindNone ObjectKind = iota
	KindInt
	KindFloat
	KindBool
	KindString
	KindList
	KindMap
	KindBuiltinFunc
	KindAlias
)

// Object is an interface for all values in golflang.
// Any values musst implement this interface.
type Object interface {
	Kind() ObjectKind // the kind of this object
	String() string
	AsBool() bool
	Equal(o Object) bool
}

/* === force implementation of Object === */

var _ Object = Int(0)
var _ Object = Float(0.0)
var _ Object = Bool(false)

// var _ Object = String("")
// var _ Object = List([]Object{})
// var _ Object = Map(make(map[Object]Object))
// var _ Object = BuiltinFunc()
// var _ Object = Alias()
