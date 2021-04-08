package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	dll := New()
	assert.NotNil(t, dll)
}

func TestIsEmpty(t *testing.T) {
	dll := New()
	assert.True(t, dll.IsEmpty())

	dll.InsertHead(1)
	assert.False(t, dll.IsEmpty())
}

func TestHeadOne(t *testing.T) {
	dll := New()
	dll.InsertHead(1)

	assert.False(t, dll.IsEmpty())
	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, dll.Head, dll.Tail)
}

func TestHeadMore(t *testing.T) {
	dll := New()
	node1 := dll.InsertHead(1)
	assert.Equal(t, 1, node1.Value)
	node2 := dll.InsertHead(2)
	assert.Equal(t, 2, node2.Value)

	assert.False(t, dll.IsEmpty())
	assert.Equal(t, 2, dll.Head.Value)
	assert.Equal(t, 1, dll.Tail.Value)
	assert.NotEqual(t, dll.Head, dll.Tail)

	assert.Equal(t, 1, dll.Head.Next.Value)
	assert.Equal(t, 2, dll.Tail.Prev.Value)
}

func TestTail(t *testing.T) {
	dll := New()
	dll.InsertTail(1)

	assert.False(t, dll.IsEmpty())
	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, dll.Head, dll.Tail)
}

func TestTailMore(t *testing.T) {
	dll := New()
	dll.InsertTail(1)
	dll.InsertTail(2)

	assert.False(t, dll.IsEmpty())
	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 2, dll.Tail.Value)
}

func TestInsertAfter(t *testing.T) {
	dll := New()

	node1 := dll.InsertHead(1)
	dll.InsertTail(2)
	dll.InsertAfter(node1, 3)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 3, dll.Head.Next.Value)
	assert.Equal(t, 2, dll.Tail.Value)
}

func TestInsertBefore(t *testing.T) {
	dll := New()

	dll.InsertHead(1)
	node2 := dll.InsertTail(2)
	dll.InsertBefore(node2, 3)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 3, dll.Head.Next.Value)
	assert.Equal(t, 2, dll.Tail.Value)
}

func TestRemoveMiddle(t *testing.T) {
	dll := New()

	dll.InsertTail(1)
	node2 := dll.InsertTail(2)
	dll.InsertTail(3)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 2, dll.Head.Next.Value)
	assert.Equal(t, 3, dll.Head.Next.Next.Value)

	dll.Remove(node2)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 3, dll.Head.Next.Value)
	assert.Nil(t, dll.Head.Next.Next)

}

func TestRemoveHead(t *testing.T) {
	dll := New()

	node1 := dll.InsertTail(1)
	dll.InsertTail(2)
	dll.InsertTail(3)

	dll.Remove(node1)

	assert.Equal(t, 2, dll.Head.Value)
	assert.Equal(t, 3, dll.Head.Next.Value)
	assert.Nil(t, dll.Head.Next.Next)
}

func TestRemoveTail(t *testing.T) {
	dll := New()

	dll.InsertTail(1)
	dll.InsertTail(2)
	node3 := dll.InsertTail(3)

	dll.Remove(node3)

	assert.Equal(t, 1, dll.Head.Value)
	assert.Equal(t, 2, dll.Head.Next.Value)
	assert.Nil(t, dll.Head.Next.Next)

}
