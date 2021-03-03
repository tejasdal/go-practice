package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
	top   int
}

func NewStack(size int) *Stack {
	return &Stack{
		items: make([]int, size),
		top: 0,
	}
}

func (stack *Stack) Push(newItem int) error {

	if stack.top >= len(stack.items) {
		return errors.New(fmt.Sprintf("stack overflow, stack size: %d and top points at %d", len(stack.items), stack.top))
	}
	stack.items[stack.top] = newItem
	stack.top++
	return nil
}

func (stack *Stack) Pop() (int,error) {

	if stack.top <= 0 {
		return -1, errors.New("Stack is empty!!")
	}
	stack.top--
	item := stack.items[stack.top]
	return item, nil
}

func (stack Stack) Print()  {
	for i, val := range stack.items {
		if i >= stack.top {
			break
		}
		fmt.Println(val)
	}
}
