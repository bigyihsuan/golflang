package obj

import (
	"fmt"
	"strings"
)

type Str string

func ZeroStr() Str {
	return Str("")
}

func NewStr(s string) Str {
	return Str(strings.Trim(s, "\""))
}

// Bool implements Obj.
func (s Str) Bool() bool {
	return s != ""
}

// Equal implements Obj.
func (s Str) Equal(o Obj) bool {
	switch o.Kind() {
	case ObjKindStr:
		return s == o.(Str)
	default:
		return false
	}
}

// Kind implements Obj.
func (s Str) Kind() ObjKind {
	return ObjKindStr
}

// Repr implements Obj.
func (s Str) Repr() string {
	return fmt.Sprintf("`%s`", strings.ReplaceAll(string(s), "\n", "\\n"))
}

// String implements Obj.
func (s Str) String() string {
	return string(s)
}

func (s Str) Hash() Hash {
	return Hash(fmt.Sprintf("%#v", s))
}
