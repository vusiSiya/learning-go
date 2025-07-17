package queue

import "reflect"

type Node[T any] struct {
	Prev  *Node[T]
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func Dequeue[T any](queue *Queue[T]) T {
	if queue.Head == nil {
		var zero T
		return zero
	}
	var headCopy = queue.Head.Value
	queue.Head = queue.Head.Next
	queue.Size--
	return headCopy
}

func Enqueue[T any](queue *Queue[T], value T) {
	if queue.Head == nil {
		var node = &Node[T]{Prev: nil, Value: value, Next: nil}
		node.Next = nil
		queue.Head = node
		queue.Tail = node
	} else {
		var node = &Node[T]{Prev: nil, Value: value, Next: nil}
		queue.Tail.Next = node
		queue.Tail = node
	}
	queue.Size++
}

func Peek[T any](queue *Queue[T]) T {
	if queue.Head != nil {
		return queue.Head.Value
	}
	var zero T
	return zero
}

func Contains[T any](queue *Queue[T], value T) bool {
	if queue.Head != nil {
		var current = queue.Head
		for current != nil {
			if reflect.DeepEqual(current.Value, value) {
				return true
			}
			current = current.Next
		}
	}
	return false
}

func IndexOf[T any](queue *Queue[T], value T) int {
	if queue.Head != nil {
		var current = queue.Head
		var index int = 0
		for current != nil {
			if reflect.DeepEqual(current.Value, value) {
				return index
			}
			current = current.Next
			index++
		}
	}
	return -1
}

func (q *Queue[T]) ToArray() []T {
	var array []T
	var i = 0
	for q.Size > 0 {
		var item = Dequeue(q)
		array = append(array, item)
		i++
	}
	return array
}
