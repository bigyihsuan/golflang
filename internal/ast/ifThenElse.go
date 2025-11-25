package ast

import "fmt"

type IfThenElse struct {
	Cond ExprList
	Then ExprList
	Else ExprList
}

func (ift IfThenElse) node() {}
func (ift IfThenElse) expr() {}
func (ift IfThenElse) String() string {
	cond := ift.Cond.String()
	then := ift.Then.String()
	else_ := ift.Else.String()
	return fmt.Sprintf("(if %s then %s else %s)", cond, then, else_)
}
