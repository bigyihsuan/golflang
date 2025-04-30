package ast

import "bigyihsuan/golflang/internal/obj"

type Ident obj.Ident

func (i Ident) expr() {}
func (i Ident) node() {}
func (i Ident) stmt() {}

func (i Ident) String() string { return obj.Ident(i).Repr() }
