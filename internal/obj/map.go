package obj

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

type Map struct {
	m map[Hash]MapEntry
}

type MapEntry struct{ K, V Obj }

func ZeroMap() Map {
	m := make(map[Hash]MapEntry)
	return Map{m}
}

func MapFromPairs(kvs ...Obj) Map {
	if len(kvs)%2 != 0 {
		panic("Map FromPairs expects even number of arguments")
	}
	out := ZeroMap()
	pairs := slices.Chunk(kvs, 2)
	for pair := range pairs {
		k, v := pair[0], pair[1]
		out.Set(k, v)
	}
	return out
}

func MapFromEntries(entries ...MapEntry) Map {
	out := ZeroMap()
	for _, e := range entries {
		out.SetEntry(e)
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
	case ObjKindMap:
		return maps.EqualFunc(m.m, o.(Map).m,
			func(e1, e2 MapEntry) bool { return e1.K.Equal(e2.K) && e1.V.Equal(e2.V) })
	default:
		return false
	}
}

// Kind implements Obj.
func (m Map) Kind() ObjKind {
	return ObjKindMap
}

// Repr implements Obj.
func (m Map) Repr() string {
	vs := []string{}
	for _, e := range m.m {
		vs = append(vs, fmt.Sprintf("%s:%s", e.K.Repr(), e.V.Repr()))
	}
	slices.Sort(vs)
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}

// String implements Obj.
func (m Map) String() string {
	vs := []string{}
	for _, e := range m.m {
		vs = append(vs, fmt.Sprintf("%s:%s", e.K.String(), e.V.String()))
	}
	slices.Sort(vs)
	return fmt.Sprintf("{%s}", strings.Join(vs, ","))
}

// Hash implements Obj.
func (m Map) Hash() Hash {
	return Hash(fmt.Sprintf("%#v", m))
}

func (m *Map) Set(k, v Obj) {
	e := MapEntry{k, v}
	m.m[e.K.Hash()] = e
}

func (m *Map) SetEntry(e MapEntry) {
	m.m[e.K.Hash()] = e
}

func (m Map) Get(k Obj) (v Obj, ok bool) {
	e, ok := m.m[k.Hash()]
	if ok {
		return e.V, ok
	}
	return nil, false
}

func (m Map) Keys() iter.Seq[Obj] {
	return func(yield func(Obj) bool) {
		for _, kv := range m.m {
			if !yield(kv.K) {
				return
			}
		}
	}
}
func (m Map) Values() iter.Seq[Obj] {
	return func(yield func(Obj) bool) {
		for _, kv := range m.m {
			if !yield(kv.V) {
				return
			}
		}
	}
}

func (m Map) Entries() iter.Seq[MapEntry] {
	return func(yield func(MapEntry) bool) {
		for _, kv := range m.m {
			if !yield(kv) {
				return
			}
		}
	}
}

func (m Map) Pairs() iter.Seq2[Obj, Obj] {
	return func(yield func(Obj, Obj) bool) {
		for _, kv := range m.m {
			if !yield(kv.K, kv.V) {
				return
			}
		}
	}
}
