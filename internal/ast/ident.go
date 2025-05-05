package ast

import (
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"strings"
)

type Ident string

func (i Ident) expr() {}
func (i Ident) node() {}
func (i Ident) stmt() {}

func (i Ident) String() string { return fmt.Sprintf("ident(%s)", string(i)) }

type IdentList []Ident

func (i IdentList) String() string {
	return fmt.Sprintf("ident(%s)", strings.Join(util.SliceMap(i, func(i Ident) string { return i.String() }), ","))
}
