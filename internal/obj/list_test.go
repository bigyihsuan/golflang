package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_ZeroList(t *testing.T) {
	assert.Equal(t, List{}, ZeroList())
}

func TestList_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    List
		exp  bool
	}{
		{"falsey empty", ZeroList(), false},
		{"truthy non-empty", List{Int(123)}, true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestList_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    List
		r    Obj
		exp  bool
	}{
		{"empty to empty truthy", List{}, List{}, true},
		{"empty to non-empty falsey", List{}, List{Str("something")}, false},
		{"non-empty to non-empty different elements falsey", List{Int(1234)}, List{Str("something")}, false},
		{"non-empty to non-empty same elements truthy", List{Str("something")}, List{Str("something")}, true},
		{"list to any falsey", List{}, Int(1234), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestList_Kind(t *testing.T) {
	v := ZeroList()
	assert.Equal(t, KindList, v.Kind(), "list should have KindList")
}

func TestList_Repr(t *testing.T) {
	v := List{Int(1), Dec(2.2), Str("333"), List{Int(4), Int(4), Int(4), Int(4)}}
	assert.Equal(t, v.Repr(), "[1,2.2,`333`,[4,4,4,4]]")
}

func TestList_String(t *testing.T) {
	v := List{Int(1), Dec(2.2), Str("333"), List{Int(4), Int(4), Int(4), Int(4)}}
	assert.Equal(t, v.String(), "[1,2.2,333,[4,4,4,4]]")
}
