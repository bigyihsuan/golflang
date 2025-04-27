package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroString(t *testing.T) {
	assert.Equal(t, Str(""), ZeroStr())
}

func TestString_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Str
		exp  bool
	}{
		{"falsey empty", Str(""), false},
		{"truthy non-empty", Str("hello world"), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestString_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Str
		r    Obj
		exp  bool
	}{
		{"string to string true", Str("hello"), Str("hello"), true},
		{"string to string false", Str("hello"), Str("world"), false},
		{"string to else false", Str("hello"), Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestString_Kind(t *testing.T) {
	v := ZeroStr()
	assert.Equal(t, KindStr, v.Kind(), "str should have KindString")
}

func TestString_Repr(t *testing.T) {
	v := Str("hello world")
	assert.Equal(t, v.Repr(), "`hello world`")
}

func TestString_String(t *testing.T) {
	v := Str("hello world")
	assert.Equal(t, v.String(), "hello world")
}
