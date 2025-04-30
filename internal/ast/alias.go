package ast

import "fmt"

type Alias struct {
	Name  Ident
	Value Expr
}

func (a Alias) node() {}
func (a Alias) stmt() {}

func (a Alias) String() string {
	return fmt.Sprintf("(%s := %s)", a.Name.String(), a.Value.String())
}
