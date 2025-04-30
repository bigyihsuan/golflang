package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"strings"
)

var _ obj.Obj = (*Lambda)(nil)

type Lambda struct {
	Args []obj.Ident
	Body ast.Expr
}

// Bool implements Obj.
func (l Lambda) Bool() bool {
	return true
}

// Equal implements Obj.
func (l Lambda) Equal(o obj.Obj) bool {
	return false
}

// Hash implements Obj.
func (l Lambda) Hash() obj.Hash {
	return obj.Hash(fmt.Sprintf("%#v", l))
}

// Kind implements Obj.
func (l Lambda) Kind() obj.ObjKind {
	return obj.ObjKindLambda
}

// Repr implements Obj.
func (l Lambda) Repr() string {
	a := strings.Join(util.SliceMap(l.Args, func(i obj.Ident) string { return i.Repr() }), ",")
	return fmt.Sprintf("\\%s => %s", a, l.Body.String())
}

// String implements Obj.
func (l Lambda) String() string {
	a := strings.Join(util.SliceMap(l.Args, func(i obj.Ident) string { return i.String() }), ",")
	return fmt.Sprintf("\\%s => %s", a, l.Body.String())
}
