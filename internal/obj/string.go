package obj

import "fmt"

type Str string

func ZeroString() Str {
	return Str("")
}

// Bool implements Obj.
func (s Str) Bool() bool {
	return s != ""
}

// Equal implements Obj.
func (s Str) Equal(o Obj) bool {
	switch o.Kind() {
	case KindString:
		return s == o.(Str)
	default:
		return false
	}
}

// Kind implements Obj.
func (s Str) Kind() ObjKind {
	return KindString
}

// Repr implements Obj.
func (s Str) Repr() string {
	return fmt.Sprintf("`%s`", s)
}

// String implements Obj.
func (s Str) String() string {
	return string(s)
}
