package stack

import "fmt"

type Stack struct {
	Head  *Node
	Tail  *Node
	Count int
}

type Node struct {
	Prev    *Node
	Element interface{}
	Next    *Node
}

func NewStack() *Stack {
	return &Stack{
		Head:  nil,
		Tail:  nil,
		Count: 0,
	}
}

func Push(stack *Stack, item interface{}) {
	if stack.Head == nil {
		var node = &Node{
			Prev:    nil,
			Element: item,
			Next:    nil,
		}
		stack.Head = node
		stack.Tail = node
	} else {
		var node = &Node{
			Prev:    stack.Tail,
			Element: item,
			Next:    nil,
		}
		stack.Tail.Next = node
		stack.Tail = node
	}
	stack.Count++
}

func Pop(stack *Stack) interface{} {
	if stack.Head == nil {
		return nil
	}
	var item = stack.Tail.Element
	var prev *Node
	var current = stack.Head //behaves as an iCounter

	for current != stack.Tail {
		prev = current
		current = current.Next
	}
	stack.Tail = prev
	stack.Count--
	return item
}

func Peek(stack *Stack) interface{} {
	if stack.Head == nil {
		return nil
	}
	return stack.Tail.Element
}

func Clear(stack *Stack) {
	stack.Head = nil
	stack.Tail = nil
	stack.Count = 0
	stack = nil
}

func Contains(stack *Stack, item interface{}) bool {
	var current = stack.Head
	for current != nil {
		if current.Element == item {
			return true
		}
		current = current.Next
	}
	return false
}

func PrintElements(stack *Stack) {
	var current = stack.Head
	for current != nil {
		fmt.Printf("\nname: %v", current.Element)
		current = current.Next
	}
}
