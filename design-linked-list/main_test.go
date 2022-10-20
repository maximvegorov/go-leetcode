package designlinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	got1 := list.Get(1)
	assert.Equal(t, 2, got1)
	list.DeleteAtIndex(1)
	got2 := list.Get(1)
	assert.Equal(t, 3, got2)
}
