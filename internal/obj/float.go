package obj

import (
	"fmt"
	"math"
)

type Float float64

func ZeroFloat() Float {
	var f Float
	return f
}

// AsBool implements Object.
func (f Float) AsBool() bool {
	return !math.IsNaN(float64(f)) && f != 0.0
}

// Equal implements Object.
func (f Float) Equal(o Obj) bool {
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
func (f Float) Kind() ObjKind {
	return KindFloat
}

// String implements Object.
func (f Float) String() string {
	return fmt.Sprint(float64(f))
}
