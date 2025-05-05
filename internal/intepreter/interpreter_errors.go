package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

type ErrMismatchedKinds struct{ L, R obj.Obj }

func (e ErrMismatchedKinds) Error() string {
	return fmt.Sprintf("mismatched kinds: %s (%s) and %s (%s)", e.L.Repr(), e.L.Kind(), e.R.Repr(), e.R.Kind())
}
