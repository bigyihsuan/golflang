package obj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_ZeroInt(t *testing.T) {
	assert.Equal(t, Int(0), ZeroInt())
}

func TestInt_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Int
		exp  bool
	}{
		{"falsey zero", Int(0), false},
		{"truthy non-zero positive", Int(5), true},
		{"truthy non-zero negative", Int(-5), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestInt_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Int
		r    Obj
		exp  bool
	}{
		{"int to int true", Int(123), Int(123), true},
		{"int to int false", Int(123), Int(-123), false},
		{"int to float true", Int(123), Float(123), true},
		{"int to float false", Int(123), Float(123.123), false},
		{"int to else false", Int(123), Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestInt_Kind(t *testing.T) {
	v := ZeroInt()
	assert.Equal(t, KindInt, v.Kind(), "int should have KindInt")
}

func TestInt_Repr(t *testing.T) {
	v := Int(123456)
	assert.Equal(t, v.Repr(), "123456")
}

func TestInt_String(t *testing.T) {
	v := Int(123456)
	assert.Equal(t, v.String(), "123456")
}
