package main

import (
	"fmt"
	"strings"
)

func (l *List[T]) String() string {
	var result strings.Builder
	currentNode := l.Head
	for currentNode != nil {
		result.WriteString(fmt.Sprintf("%v ->", currentNode.Value))
		currentNode = currentNode.Next
	}
	return result.String()
}

type Equatable[T any] interface {
	Equals(other Equatable[T]) bool
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) InsertFront(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  l.Head,
	}

	l.Head = newNode
}

func (l *List[T]) InsertBack(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil { // The list is empty
		l.Head = newNode // The new node becomes both the head and the tail
		l.Tail = newNode

	} else { // The list is not empty
		l.Tail.Next = newNode // Link the current tail node to the new node
		l.Tail = newNode      // Update the tail to the new node
	}

}

func (l *List[T]) InsertAt(index int, value T) error {
	if index == 0 {
		l.InsertFront(value)
		return nil
	}

	// Get the node just before the position where the new node should be inserted
	previousNode, err := l.Get(index - 1)
	if err != nil {
		return err // This will handle out-of-bounds by index - 1
	}

	// Create the new node and link it correctly
	newNode := &Node[T]{
		Value: value,
		Next:  previousNode.Next, // New node should point to the current node at the index
	}

	// Insert the new node into the list
	previousNode.Next = newNode

	// Special case: if inserting as the last node
	if newNode.Next == nil {
		l.Tail = newNode
	}

	return nil
}

func (l *List[T]) RemoveFront() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}

	if l.Head == l.Tail {
		l.Tail = nil
		l.Head = nil

	} else {
		l.Head = l.Head.Next
	}

	return nil
}

func (l *List[T]) RemoveBack() error {
	if l.Tail == nil {
		return fmt.Errorf("list is empty")
	}

	if l.Head == l.Tail {
		l.Tail = nil
		l.Head = nil

	} else {
		currentNode := l.Head
		for currentNode.Next != l.Tail {
			currentNode = currentNode.Next
		}
		currentNode.Next = nil
		l.Tail = currentNode
	}

	return nil
}

func (l *List[T]) RemoveAt(index int) error {
	if index == 0 {
		l.RemoveFront()
		return nil
	}

	currentNodeAtIndex, err := l.Get(index)

	if err != nil {
		return err
	}

	if currentNodeAtIndex == l.Tail {
		l.RemoveBack()
		return nil
	}

	previousNode, err := l.Get(index - 1)

	if err != nil {
		return err
	}

	previousNode.Next = currentNodeAtIndex.Next

	return nil
}

func (l *List[T]) Get(index int) (*Node[T], error) {
	currentIndex := 0
	currentNode := l.Head
	for currentNode != nil {
		if currentIndex == index {
			return currentNode, nil
		}
		currentIndex++
		currentNode = currentNode.Next
	}

	var zeroValue *Node[T] // Prepare a zero value of T to return in case of error
	return zeroValue, fmt.Errorf("index out of bounds")
}

// Find searches for the first node containing a value equal to the provided value.
func (l *List[T]) Find(value T) (*Node[T], error) {
	if equatableVal, ok := any(value).(Equatable[T]); ok {
		currentNode := l.Head
		for currentNode != nil {
			if nodeVal, ok := any(currentNode.Value).(Equatable[T]); ok {
				if nodeVal.Equals(equatableVal) {
					return currentNode, nil
				}
			}
			currentNode = currentNode.Next
		}
	}
	return nil, fmt.Errorf("value not found")
}

func mainGenericTypes() {
	// Create a new list of integers
	list := &List[int]{}

	// Insert some integers
	list.InsertBack(1)
	list.InsertBack(2)
	list.InsertBack(3)

	fmt.Printf("%s\n", list)

	// Insert at the front
	list.InsertFront(0)
	fmt.Printf("%s\n", list)

	// Insert at index 2
	list.InsertAt(2, 1)
	fmt.Printf("%s\n", list)

	// Remove the last element
	list.RemoveBack()
	fmt.Printf("%s\n", list)

	// Remove the first element
	list.RemoveFront()
	fmt.Printf("%s\n", list)

	// Remove the element at index 1
	list.RemoveAt(1)
	fmt.Printf("%s\n", list)

	// Find an element
	node, err := list.Find(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(node.Value)
	}

	// Print the list
	fmt.Printf("%s\n", list)

}
