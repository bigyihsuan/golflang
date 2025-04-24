package obj

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Map struct {
	m    map[Obj]Obj
	keys []Obj
}

func ZeroMap() Map {
	return Map{m: make(map[Obj]Obj), keys: []Obj{}}
}

func MapFromMap(m map[Obj]Obj) Map {
	out := ZeroMap()
	out.m = m
	keys := slices.Collect(maps.Keys(m))
	if keys == nil {
		keys = []Obj{}
	}
	out.keys = keys
	return out
}

func MapFromPairs(kvs ...Obj) Map {
	if len(kvs)%2 != 0 {
		panic("Map FromPairs expects even number of arguments")
	}
	out := ZeroMap()
	pairs := slices.Chunk(kvs, 2)
	for pair := range pairs {
		k, v := pair[0], pair[1]
		out.Add(k, v)
	}
	return out
}

// Bool implements Obj.
func (m Map) Bool() bool {
	return len(m.m) > 0
}

// Equal implements Obj.
func (m Map) Equal(o Obj) bool {
	switch o.Kind() {
	case KindMap:
		return maps.EqualFunc(m.m, o.(Map).m, func(l, r Obj) bool { return l.Equal(r) })
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
	for _, k := range m.keys {
		v := m.m[k]
		vs = append(vs, fmt.Sprintf("%s:%s", k.Repr(), v.Repr()))
	}
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}

// String implements Obj.
func (m Map) String() string {
	vs := []string{}
	for _, k := range m.keys {
		v := m.m[k]
		vs = append(vs, fmt.Sprintf("%s:%s", k.String(), v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}

// replicates m[k] = v
func (m *Map) Add(k, v Obj) {
	m.m[k] = v
	if !slices.ContainsFunc(m.keys, func(o Obj) bool { return o.Equal(k) }) {
		m.keys = append(m.keys, k)
	}
}

// replicates v,ok := m[k]
func (m *Map) Get(k Obj) (Obj, bool) {
	v, ok := m.m[k]
	return v, ok
}

// replicates delete(m,k)
func (m *Map) Delete(k Obj) {
	_, ok := m.Get(k)
	if ok {
		delete(m.m, k)
		i := slices.IndexFunc(m.keys, func(o Obj) bool { return o.Equal(k) })
		if i < 0 {
			return
		}
		m.keys = slices.Delete(m.keys, i, i+1)
	}
}
