package obj

import (
	"fmt"
	"math"
)

type Dec float64

func ZeroFloat() Dec {
	var f Dec
	return f
}

// Bool implements Object.
func (f Dec) Bool() bool {
	return !math.IsNaN(float64(f)) && f != 0.0
}

// Equal implements Object.
func (f Dec) Equal(o Obj) bool {
	switch o.Kind() {
	case KindFloat:
		return f == o.(Dec)
	case KindInt:
		return f == Dec(o.(Int))
	default:
		return false
	}
}

// Kind implements Object.
func (f Dec) Kind() ObjKind {
	return KindFloat
}

// Repr implements Obj.
func (f Dec) Repr() string {
	return f.String()
}

// String implements Object.
func (f Dec) String() string {
	return fmt.Sprint(float64(f))
}
