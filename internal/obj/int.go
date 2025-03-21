package obj

import "strconv"

type Int int64

// AsBool implements Object.
func (i Int) AsBool() bool {
	return i != 0
}

// Equal implements Object.
func (i Int) Equal(o Object) bool {
	switch o.Kind() {
	case KindInt:
		return i == o.(Int)
	case KindFloat:
		return Float(i) == o.(Float)
	default:
		return false
	}
}

// Kind implements Object.
func (i Int) Kind() ObjectKind {
	return KindInt
}

// String implements Object.
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}
