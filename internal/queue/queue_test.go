package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Len(t *testing.T) {
	q := New(1, 2, 3)
	exp := 3
	act := q.Len()
	assert.Equal(t, exp, act)
}

func TestQueue_Peek_filled(t *testing.T) {
	q := New(1, 2, 3)
	exp := 1
	act, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestQueue_Peek_empty(t *testing.T) {
	q := New[int]()
	_, ok := q.Peek()
	assert.False(t, ok)
}

func TestQueue_Enqueue(t *testing.T) {
	q := New(1, 2, 3)
	q.Enqueue(4)
	exp := New(1, 2, 3, 4)
	assert.Equal(t, exp, q)
}

func TestQueue_Dequeue_filled(t *testing.T) {
	q := New(1, 2, 3)
	q.Dequeue()
	q.Dequeue()
	act, ok := q.Dequeue()
	exp := 3
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestQueue_Dequeue_single(t *testing.T) {
	q := New(1)
	exp := 1
	act, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

func TestQueue_Dequeue_empty(t *testing.T) {
	q := New[int]()
	_, ok := q.Dequeue()
	assert.False(t, ok)
}
