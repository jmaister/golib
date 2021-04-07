package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New()
	assert.NotNil(t, l, "New list not created properly")

}

func TestIsEmpty(t *testing.T) {
	l := New()
	assert.True(t, l.IsEmpty(), "List should be empty")
}

func TestAppendOne(t *testing.T) {
	l := New()
	l.Append(9)

	assert.False(t, l.IsEmpty(), "List not should be empty")
	assert.Equal(t, 9, l.Head.Value)
	assert.Equal(t, 9, l.Tail.Value)
	assert.Nil(t, l.Head.Next)
	assert.Nil(t, l.Tail.Next)
}

func TestAppendTwo(t *testing.T) {
	l := New()
	l.Append(8)
	l.Append(9)

	assert.False(t, l.IsEmpty(), "List not should be empty")
	assert.Equal(t, 8, l.Head.Value)
	assert.Equal(t, 9, l.Tail.Value)
	assert.NotNil(t, l.Head.Next)
	assert.Nil(t, l.Tail.Next)
	assert.Equal(t, l.Head.Next, l.Tail)
}
