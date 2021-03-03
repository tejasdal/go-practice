package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {

	var myStack = NewStack(5)
	myStack.Push(2)
	myStack.Push(3)
	assert.Equal(t, myStack.top, 2)
}

func TestStack_PushStackOverflow(t *testing.T) {

	var myStack = NewStack(1)
	myStack.Push(2)
	assert.Error(t, myStack.Push(3))
}

func TestStack_Pop(t *testing.T) {

	var myStack = NewStack(3)
	myStack.Push(2)
	myStack.Push(3)
	topItem, _ := myStack.Pop()
	assert.Equal(t, topItem, 3)
	assert.Equal(t, myStack.top, 1)
}

func TestStack_PopUnderflow(t *testing.T) {
	var myStack = NewStack(1)
	myStack.Push(2)
	myStack.Pop()
	_, err := myStack.Pop()
	assert.Error(t, err)
}