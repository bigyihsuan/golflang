package ast

import "bigyihsuan/golflang/internal/obj"

type Ident struct {
	Name obj.Ident
}

func (i Ident) expr()          {}
func (i Ident) node()          {}
func (i Ident) String() string { return i.Name.Repr() }
