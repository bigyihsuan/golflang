package obj

import (
	"fmt"
)

type Float float64

func ZeroFloat() Float {
	var f Float
	return f
}

// AsBool implements Object.
func (f Float) AsBool() bool {
	return f != 0.0
}

// Equal implements Object.
func (f Float) Equal(o Object) bool {
	switch o.Kind() {
	case KindFloat:
		return f == o.(Float)
	case KindInt:
		return f == Float(o.(Int))
	default:
		return false
	}
}

// Kind implements Object.
func (f Float) Kind() ObjectKind {
	return KindFloat
}

// String implements Object.
func (f Float) String() string {
	return fmt.Sprint(float64(f))
}
