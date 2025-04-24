package obj

import "strconv"

type Bool bool

func ZeroBool() Bool {
	var b Bool
	return b
}

// AsBool implements Object.
func (b Bool) AsBool() bool {
	return bool(b)
}

// Equal implements Object.
func (b Bool) Equal(o Obj) bool {
	switch o.Kind() {
	case KindBool:
		return b == o.(Bool)
	default:
		return false
	}
}

// Kind implements Object.
func (b Bool) Kind() ObjKind {
	return KindBool
}

// String implements Object.
func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}
