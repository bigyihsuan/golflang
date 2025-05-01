package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Len(t *testing.T) {
	s := New(1, 2, 3)
	exp := 3
	act := s.Len()
	assert.Equal(t, exp, act)
}

func TestStack_Peek_filled(t *testing.T) {
	s := New(1, 2, 3)
	exp := 1
	act, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestStack_Peek_empty(t *testing.T) {
	s := New[int]()
	_, ok := s.Peek()
	assert.False(t, ok)
}

func TestStack_Push(t *testing.T) {
	s := New(1, 2, 3)
	s.Push(4)
	exp := New(1, 2, 3, 4)
	assert.Equal(t, exp, s)
}

func TestStack_Pop_filled(t *testing.T) {
	s := New(1, 2, 3)
	s.Pop()
	s.Pop()
	act, ok := s.Pop()
	exp := 1
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestStack_Pop_single(t *testing.T) {
	s := New(1)
	exp := 1
	act, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestStack_Pop_empty(t *testing.T) {
	s := New[int]()
	_, ok := s.Pop()
	assert.False(t, ok)
}
