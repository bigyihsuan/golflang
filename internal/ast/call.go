package ast

import (
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"strings"
)

type Call struct {
	Name Ident
	Args []Expr
}

func (c Call) node() {}
func (c Call) stmt() {}
func (c Call) expr() {}

func (c Call) String() string {
	es := util.SliceMap(c.Args, func(e Expr) string { return e.String() })
	return fmt.Sprintf("call(%s %s)", c.Name, strings.Join(es, " "))
}
