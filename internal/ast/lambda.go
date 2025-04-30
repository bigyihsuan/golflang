package ast

import "fmt"

type Lambda struct {
	Args IdentList
	Body Expr
}

func (l Lambda) node() {}
func (l Lambda) expr() {}

func (l Lambda) String() string {
	args := l.Args.String()
	body := l.Body.String()
	return fmt.Sprintf("(\\%s => %s)", args, body)
}
