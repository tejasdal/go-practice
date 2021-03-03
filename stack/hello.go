package main

import "fmt"

func main() {
	var myStack = NewStack(10)
	myStack.Push(5)
	myStack.Push(15)
	myStack.Push(25)
	myStack.Push(35)
	fmt.Println("After push:")
	myStack.Print()
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	fmt.Println("After pop:")
	myStack.Print()
}
