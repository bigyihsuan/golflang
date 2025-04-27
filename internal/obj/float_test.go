package obj

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat_ZeroFloat(t *testing.T) {
	assert.Equal(t, Dec(0.0), ZeroDec())
}

func TestFloat_Bool(t *testing.T) {
	ts := []struct {
		desc string
		v    Dec
		exp  bool
	}{
		{"falsey zero", Dec(0), false},
		{"falsey NaN", Dec(math.NaN()), false},
		{"truthy non-zero positive", Dec(5.6), true},
		{"truthy non-zero negative", Dec(-5.6), true},
	}

	for _, test := range ts {
		act := test.v.Bool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    Dec
		r    Obj
		exp  bool
	}{
		{"float to float true", Dec(123.123), Dec(123.123), true},
		{"float to float false", Dec(123.123), Dec(-123.123), false},
		{"float to int true", Dec(123), Int(123), true},
		{"float to int false", Dec(123.123), Int(123), false},
		{"float to else false", Dec(123.123), Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_Kind(t *testing.T) {
	v := ZeroDec()
	assert.Equal(t, ObjKindDec, v.Kind(), "float should have KindFloat")
}

func TestFloat_Repr(t *testing.T) {
	v := Dec(123.456)
	assert.Equal(t, v.Repr(), "123.456")
}

func TestFloat_String(t *testing.T) {
	v := Dec(123.456)
	assert.Equal(t, v.String(), "123.456")
}
