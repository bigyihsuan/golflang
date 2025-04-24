package obj

import "fmt"

type String string

func ZeroString() String {
	return String("")
}

// Bool implements Obj.
func (s String) Bool() bool {
	return s != ""
}

// Equal implements Obj.
func (s String) Equal(o Obj) bool {
	switch o.Kind() {
	case KindString:
		return s == o.(String)
	default:
		return false
	}
}

// Kind implements Obj.
func (s String) Kind() ObjKind {
	return KindString
}

// Repr implements Obj.
func (s String) Repr() string {
	return fmt.Sprintf("`%s`", s)
}

// String implements Obj.
func (s String) String() string {
	return string(s)
}
