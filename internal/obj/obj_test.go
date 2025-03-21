package obj_test

import (
	o "bigyihsuan/golflang/internal/obj"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_AsBool(t *testing.T) {
	ts := []struct {
		desc string
		v    o.Int
		exp  bool
	}{
		{"falsey zero", o.Int(0), false},
		{"truthy non-zero positive", o.Int(5), true},
		{"truthy non-zero negative", o.Int(-5), true},
	}

	for _, test := range ts {
		act := test.v.AsBool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestInt_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    o.Int
		r    o.Object
		exp  bool
	}{
		{"int to int true", o.Int(123), o.Int(123), true},
		{"int to int false", o.Int(123), o.Int(-123), false},
		{"int to float true", o.Int(123), o.Float(123), true},
		{"int to float false", o.Int(123), o.Float(123.123), false},
		{"int to else false", o.Int(123), o.Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_AsBool(t *testing.T) {
	ts := []struct {
		desc string
		v    o.Float
		exp  bool
	}{
		{"falsey zero", o.Float(0), false},
		{"truthy non-zero positive", o.Float(5.6), true},
		{"truthy non-zero negative", o.Float(-5.6), true},
	}

	for _, test := range ts {
		act := test.v.AsBool()
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestFloat_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    o.Float
		r    o.Object
		exp  bool
	}{
		{"float to float true", o.Float(123.123), o.Float(123.123), true},
		{"float to float false", o.Float(123.123), o.Float(-123.123), false},
		{"float to int true", o.Float(123), o.Int(123), true},
		{"float to int false", o.Float(123.123), o.Int(123), false},
		{"float to else false", o.Float(123.123), o.Bool(false), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}

func TestBool_Equal(t *testing.T) {
	ts := []struct {
		desc string
		l    o.Bool
		r    o.Object
		exp  bool
	}{
		{"bool to bool true", o.Bool(false), o.Bool(false), true},
		{"bool to bool false", o.Bool(false), o.Bool(true), false},
		{"bool to else false", o.Bool(false), o.Int(123), false},
	}

	for _, test := range ts {
		act := test.l.Equal(test.r)
		assert.Equal(t, test.exp, act, test.desc)
	}
}
