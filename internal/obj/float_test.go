package obj

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat_ZeroFloat(t *testing.T) {
	assert.Equal(t, Float(0.0), ZeroFloat())
}

func TestFloat_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Float
		exp  bool
	}{
		{"falsey zero", Float(0), false},
		{"falsey NaN", Float(math.NaN()), false},
		{"truthy non-zero positive", Float(5.6), true},
		{"truthy non-zero negative", Float(-5.6), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Float
		r    Obj
		exp  bool
	}{
		{"float to float true", Float(123.123), Float(123.123), true},
		{"float to float false", Float(123.123), Float(-123.123), false},
		{"float to int true", Float(123), Int(123), true},
		{"float to int false", Float(123.123), Int(123), false},
		{"float to else false", Float(123.123), Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_Kind(t *testing.T) {
	v := ZeroFloat()
	assert.Equal(t, KindFloat, v.Kind(), "float should have KindFloat")
}

func TestFloat_Repr(t *testing.T) {
	v := Float(123.456)
	assert.Equal(t, v.Repr(), "123.456")
}

func TestFloat_String(t *testing.T) {
	v := Float(123.456)
	assert.Equal(t, v.String(), "123.456")
}
