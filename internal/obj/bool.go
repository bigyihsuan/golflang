package obj

import (
	"fmt"
	"strconv"
)

type Bool bool

func ZeroBool() Bool {
	var b Bool
	return b
}

// Bool implements Object.
func (b Bool) Bool() bool {
	return bool(b)
}

// Equal implements Object.
func (b Bool) Equal(o Obj) bool {
	switch o.Kind() {
	case ObjKindBool:
		return b == o.(Bool)
	default:
		return false
	}
}

// Kind implements Object.
func (b Bool) Kind() ObjKind {
	return ObjKindBool
}

// Repr implements Obj.
func (b Bool) Repr() string {
	return b.String()
}

// String implements Object.
func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

func (b Bool) Hash() Hash {
	return Hash(fmt.Sprintf("%#v", b))
}
