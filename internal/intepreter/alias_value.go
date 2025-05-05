package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

var _ obj.Obj = AliasValue{}

type AliasValue struct {
	ast.ExprList
}

func NewAliasValue(exprList ast.ExprList) AliasValue {
	return AliasValue{ExprList: exprList}
}

// Bool implements obj.Obj.
func (a AliasValue) Bool() bool {
	return false
}

// Equal implements obj.Obj.
func (a AliasValue) Equal(o obj.Obj) bool {
	return false
}

// Hash implements obj.Obj.
func (a AliasValue) Hash() obj.Hash {
	return obj.Hash(fmt.Sprintf("%#v", a))
}

// Kind implements obj.Obj.
func (a AliasValue) Kind() obj.ObjKind {
	return obj.ObjKindAliasValue
}

// Repr implements obj.Obj.
func (a AliasValue) Repr() string {
	return fmt.Sprintf("aliasValue(%s)", a.ExprList.String())
}

// String implements obj.Obj.
func (a AliasValue) String() string {
	return a.ExprList.String()
}
