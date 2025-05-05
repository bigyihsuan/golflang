package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
	"reflect"
)

var _ obj.Obj = (*BuiltinFunc)(nil)

type BuiltinFunc func(*Interpreter) (obj.Obj, error)

// Bool implements obj.Obj.
func (b BuiltinFunc) Bool() bool {
	return false
}

// Equal implements obj.Obj.
func (b BuiltinFunc) Equal(o obj.Obj) bool {
	return false
}

// Hash implements obj.Obj.
func (b BuiltinFunc) Hash() obj.Hash {
	return obj.Hash(fmt.Sprintf("%#v", b))
}

// Kind implements obj.Obj.
func (b BuiltinFunc) Kind() obj.ObjKind {
	return obj.ObjKindBuiltinFunc
}

// Repr implements obj.Obj.
func (b BuiltinFunc) Repr() string {
	return reflect.ValueOf(b).String()
}

// String implements obj.Obj.
func (b BuiltinFunc) String() string {
	return reflect.ValueOf(b).String()
}
