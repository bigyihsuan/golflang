package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_ZeroMap(t *testing.T) {
	assert.Equal(t, Map{m: make(map[Hash]MapEntry)}, ZeroMap())
}

func TestMap_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Map
		exp  bool
	}{
		{"falsey empty", ZeroMap(), false},
		{"truthy non-empty", MapFromPairs(Str("a"), Int(123)), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestMap_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Map
		r    Obj
		exp  bool
	}{
		{"empty to empty truthy", ZeroMap(), ZeroMap(), true},
		{"empty to non-empty falsey", ZeroMap(), MapFromPairs(Int(123), Str("something")), false},
		{"non-empty to non-empty different elements falsey", MapFromPairs(
			Int(123), Str("something"),
		), MapFromPairs(
			Dec(123.456), Str("something"),
		), false},
		{"non-empty to non-empty same elements truthy", MapFromPairs(
			Bool(true), List{Int(1), Int(2)},
			Int(123), Str("something"),
		), MapFromPairs(
			Int(123), Str("something"),
			Bool(true), List{Int(1), Int(2)},
		), true},
		{"map to any falsey", ZeroMap(), Int(1234), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestMap_Kind(t *testing.T) {
	v := ZeroMap()
	assert.Equal(t, ObjKindMap, v.Kind(), "map should have KindMap")
}

func TestMap_Repr(t *testing.T) {
	v := MapFromPairs(Int(1), Dec(2.2), Str("333"), List{Int(4), Int(4), Int(4), Int(4)})
	assert.Equal(t, v.Repr(), "{1:2.2,`333`:[4,4,4,4]}")
}

func TestMap_String(t *testing.T) {
	v := MapFromPairs(Int(1), Dec(2.2), Str("333"), List{Int(4), Int(4), Int(4), Int(4)})
	assert.Equal(t, v.String(), "{1:2.2,333:[4,4,4,4]}")
}
