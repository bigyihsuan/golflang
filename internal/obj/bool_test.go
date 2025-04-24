package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_ZeroBool(t *testing.T) {
	assert.Equal(t, Bool(false), ZeroBool())
}

func TestBool_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Bool
		exp  bool
	}{
		{"falsey false", Bool(false), false},
		{"truthy true", Bool(true), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestBool_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Bool
		r    Obj
		exp  bool
	}{
		{"bool to bool true", Bool(false), Bool(false), true},
		{"bool to bool false", Bool(false), Bool(true), false},
		{"bool to else false", Bool(false), Int(123), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestBool_Kind(t *testing.T) {
	v := ZeroBool()
	assert.Equal(t, KindBool, v.Kind(), "bool should have KindBool")
}

func TestBool_Repr(t *testing.T) {
	v := Bool(true)
	assert.Equal(t, v.Repr(), "true")
}

func TestBool_String(t *testing.T) {
	v := Bool(true)
	assert.Equal(t, v.String(), "true")
}
