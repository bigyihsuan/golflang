package ast

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
	"strings"
)

type LiteralPrimitive struct {
	Value obj.Obj
}

func (l LiteralPrimitive) node()          {}
func (l LiteralPrimitive) stmt()          {}
func (l LiteralPrimitive) expr()          {}
func (l LiteralPrimitive) lit()           {}
func (l LiteralPrimitive) String() string { return l.Value.Repr() }

type LiteralList struct {
	Value []Expr
}

func (l LiteralList) node() {}
func (l LiteralList) stmt() {}
func (l LiteralList) expr() {}
func (l LiteralList) lit()  {}
func (l LiteralList) String() string {
	es := []string{}
	for _, e := range l.Value {
		es = append(es, e.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(es, ","))
}

type LiteralMap struct {
	Value []LiteralMapEntry
}

func (l LiteralMap) node() {}
func (l LiteralMap) stmt() {}
func (l LiteralMap) expr() {}
func (l LiteralMap) lit()  {}
func (l LiteralMap) String() string {
	es := []string{}
	for _, e := range l.Value {
		s := fmt.Sprintf("%s:%s", e.K.String(), e.V.String())
		es = append(es, s)
	}
	return fmt.Sprintf("{%s}", strings.Join(es, ","))
}

type LiteralMapEntry struct {
	K, V Expr
}
