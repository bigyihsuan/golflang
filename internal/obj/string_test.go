package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroString(t *testing.T) {
	assert.Equal(t, String(""), ZeroString())
}

func TestString_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    String
		exp  bool
	}{
		{"falsey empty", String(""), false},
		{"truthy non-empty", String("hello world"), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestString_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    String
		r    Obj
		exp  bool
	}{
		{"string to string true", String("hello"), String("hello"), true},
		{"string to string false", String("hello"), String("world"), false},
		{"string to else false", String("hello"), Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestString_Kind(t *testing.T) {
	v := ZeroString()
	assert.Equal(t, KindString, v.Kind(), "str should have KindString")
}

func TestString_Repr(t *testing.T) {
	v := String("hello world")
	assert.Equal(t, v.Repr(), "`hello world`")
}

func TestString_String(t *testing.T) {
	v := String("hello world")
	assert.Equal(t, v.String(), "hello world")
}
