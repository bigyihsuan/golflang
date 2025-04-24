package obj

import "strconv"

type Int int64

func ZeroInt() Int {
	var i Int
	return i
}

// Bool implements Object.
func (i Int) Bool() bool {
	return i != 0
}

// Equal implements Object.
func (i Int) Equal(o Obj) bool {
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
func (i Int) Kind() ObjKind {
	return KindInt
}

// Repr implements Obj.
func (i Int) Repr() string {
	return i.String()
}

// String implements Object.
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}
