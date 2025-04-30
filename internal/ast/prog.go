package ast

import "strings"

type Prog struct {
	Stmts []Stmt
}

func (p Prog) node() {}
func (p Prog) String() string {
	s := []string{}
	for _, stmt := range p.Stmts {
		s = append(s, stmt.String()+";")
	}
	return strings.Join(s, " ")
}
