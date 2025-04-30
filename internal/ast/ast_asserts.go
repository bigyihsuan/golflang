package ast

var _ Node = (*Prog)(nil)
var _ Lit = (*LiteralPrimitive)(nil)
var _ Lit = (*LiteralList)(nil)
var _ Lit = (*LiteralMap)(nil)
var _ Stmt = (*LiteralPrimitive)(nil)
var _ Stmt = (*LiteralList)(nil)
var _ Stmt = (*LiteralMap)(nil)
var _ Stmt = (*Alias)(nil)
var _ Expr = (*Ident)(nil)
