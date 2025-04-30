package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

// Scope contains local aliases.
// One is created with each call of a lambda, and destroyed when the lambda finishes.
type Scope struct {
	parent  *Scope                // parent of this scope
	aliases map[obj.Ident]obj.Obj // aliases declared in this scope
}

func NewBaseScope() Scope {
	return NewScope(nil)
}

func NewScope(parent *Scope) Scope {
	return Scope{
		parent:  parent,
		aliases: make(map[obj.Ident]obj.Obj),
	}
}

func (s *Scope) SetAlias(name obj.Ident, value obj.Obj) {
	s.aliases[name] = value
}

func (s Scope) GetAlias(i obj.Ident) (value obj.Obj) {
	value, _ = s.getAlias(i)
	return value
}

func (s Scope) getAlias(name obj.Ident) (value obj.Obj, ok bool) {
	if value, inThisScope := s.aliases[name]; inThisScope {
		return value, true
	} else if value, inParentScope := s.parent.getAlias(name); inParentScope {
		return value, true
	} else {
		panic(fmt.Errorf(InterpreterErrorAliasNotFound, name.String()))
	}
}

func (s Scope) String() string {
	parent := "<base>"
	if s.parent != nil {
		parent = s.parent.String()
	}
	aliases := fmt.Sprint(s.aliases)
	return fmt.Sprintf("{%s <- %s}", parent, aliases)
}
