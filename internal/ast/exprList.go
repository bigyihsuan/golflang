package ast

import (
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"strings"
)

type ExprList []Expr

func (e ExprList) node() {}
func (e ExprList) stmt() {}
func (e ExprList) expr() {}

func (e ExprList) String() string {
	return fmt.Sprintf("(%s)",
		strings.Join(util.SliceMap[Expr, string](e, func(e Expr) string { return e.String() }), " "))
}
