package scope

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

// Scope contains local aliases.
// One is created with each call of a lambda, and destroyed when the lambda finishes.
type Scope struct {
	Parent  *Scope                // parent of this scope
	aliases map[obj.Ident]obj.Obj // aliases declared in this scope
}

func NewBase() Scope {
	return New(nil)
}

func New(parent *Scope) Scope {
	return Scope{
		Parent:  parent,
		aliases: make(map[obj.Ident]obj.Obj),
	}
}

func (s *Scope) SetAlias(name obj.Ident, value obj.Obj) {
	s.aliases[name] = value
}

func (s Scope) GetAlias(i obj.Ident) (value obj.Obj, err error) {
	value, ok := s.getAlias(i)
	if !ok {
		return value, ErrUnknownAlias{i.String()}
	}
	return value, nil
}

func (s Scope) getAlias(name obj.Ident) (value obj.Obj, ok bool) {
	if value, inThisScope := s.aliases[name]; inThisScope {
		return value, true
	} else if value, inParentScope := s.Parent.getAlias(name); inParentScope {
		return value, true
	} else {
		return nil, false
	}
}

func (s Scope) String() string {
	parent := "<base>"
	if s.Parent != nil {
		parent = s.Parent.String()
	}
	aliases := fmt.Sprint(s.aliases)
	return fmt.Sprintf("{%s <- %s}", parent, aliases)
}
