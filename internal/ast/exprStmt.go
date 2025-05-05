package ast

import "fmt"

type ExprStmt struct {
	ExprList
}

func (e ExprStmt) node() {}
func (e ExprStmt) stmt() {}

func (e ExprStmt) String() string {
	return fmt.Sprintf("exprStmt(%s)", e.ExprList.String())
}
