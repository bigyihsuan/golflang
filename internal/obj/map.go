package obj

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Map map[Obj]Obj

func ZeroMap() Map {
	return Map(make(map[Obj]Obj))
}

func MapFromMap(m map[Obj]Obj) Map {
	return Map(m)
}

func MapFromPairs(kvs ...Obj) Map {
	if len(kvs)%2 != 0 {
		panic("Map FromPairs expects even number of arguments")
	}
	out := ZeroMap()
	pairs := slices.Chunk(kvs, 2)
	for pair := range pairs {
		k, v := pair[0], pair[1]
		out[k] = v
	}
	return out
}

// Bool implements Obj.
func (m Map) Bool() bool {
	return len(m) > 0
}

// Equal implements Obj.
func (m Map) Equal(o Obj) bool {
	switch o.Kind() {
	case KindMap:
		return maps.EqualFunc(m, o.(Map), func(l, r Obj) bool { return l.Equal(r) })
	default:
		return false
	}
}

// Kind implements Obj.
func (m Map) Kind() ObjKind {
	return KindMap
}

// Repr implements Obj.
func (m Map) Repr() string {
	vs := []string{}
	for k, v := range m {
		vs = append(vs, fmt.Sprintf("%s:%s", k.Repr(), v.Repr()))
	}
	slices.Sort(vs)
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}

// String implements Obj.
func (m Map) String() string {
	vs := []string{}
	for k, v := range m {
		vs = append(vs, fmt.Sprintf("%s:%s", k.String(), v.String()))
	}
	slices.Sort(vs)
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}
