package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"strings"
)

type ErrMismatchedKinds struct{ L, R obj.Obj }

func (e ErrMismatchedKinds) Error() string {
	return fmt.Sprintf("mismatched kinds: %s (%s) and %s (%s)", e.L.Repr(), e.L.Kind(), e.R.Repr(), e.R.Kind())
}

type ArgAcceptedTypes map[obj.Hash][]obj.ObjKind

type ErrIncorrectArgumentTypes struct {
	Name string
	Args ArgAcceptedTypes // key is the argument, and value is the acceptable types
}

func (e ErrIncorrectArgumentTypes) Error() string {
	s := []string{}
	for o, t := range e.Args {
		s = append(
			s,
			fmt.Sprintf(
				"got %s, want %s;",
				o,
				strings.Join(
					util.SliceMap(
						t, func(o obj.ObjKind) string { return o.String() }),
					"/")))
	}
	return fmt.Sprintf("incorrect argument types for %s: %s", e.Name, strings.Join(s, " "))
}
