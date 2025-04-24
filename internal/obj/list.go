package obj

import (
	"bigyihsuan/golflang/internal/slice"
	"fmt"
	"slices"
	"strings"
)

type List []Obj

func ZeroList() List {
	return List{}
}

// Bool implements Obj.
func (l List) Bool() bool {
	return len(l) > 0
}

// Equal implements Obj.
func (l List) Equal(o Obj) bool {
	switch o.Kind() {
	case KindList:
		o := o.(List)
		return slices.EqualFunc(l, o, func(l, r Obj) bool { return l.Equal(r) })
	default:
		return false
	}
}

// Kind implements Obj.
func (l List) Kind() ObjKind {
	return KindList
}

// Repr implements Obj.
func (l List) Repr() string {
	return fmt.Sprintf("[%s]", strings.Join(slice.Map(l, func(o Obj) string { return o.Repr() }), ","))
}

// String implements Obj.
func (l List) String() string {
	return fmt.Sprintf("[%s]", strings.Join(slice.Map(l, func(o Obj) string { return o.String() }), ","))
}
