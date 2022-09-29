package datastructures

import (
	"fmt"
	"strings"
)

type DllNode[T comparable] struct {
	Value T
	Prev  *DllNode[T]
	Next  *DllNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Length int
	head   *DllNode[T]
	tail   *DllNode[T]
}

func (this DllNode[T]) String() string {
	return fmt.Sprintf("%v", this.Value)
}

func (this DoublyLinkedList[T]) String() string {
	sb := strings.Builder{}

	for it := this.head; it != nil; it = it.Next {
		sb.WriteString(fmt.Sprintf("%v ", it))
	}

	return sb.String()
}

func (this *DoublyLinkedList[T]) Prepend(item T) {
	node := DllNode[T]{
		Value: item,
	}

	if this.head == nil {
		this.head = this.tail
		this.tail = &node
		return
	}

	node.Next = this.head
	this.head.Prev = &node
	this.head = &node
}

func (this *DoublyLinkedList[T]) InsertAt(item T, idx int) {
	if idx > this.Length {
		panic("oh no")
	}

	if idx == this.Length {
		this.Append(item)
		return
	} else if idx == 0 {
		this.Prepend(item)
		return
	}

	this.Length += 1
	curr := this.getAt(idx)

	node := DllNode[T]{
		Value: item,
	}

	node.Next = curr
	node.Prev = curr.Prev
	curr.Prev = &node
	if node.Prev != nil {
		curr.Prev.Next = curr
	}
}

func (this *DoublyLinkedList[T]) Append(item T) {
	this.Length += 1
	node := DllNode[T]{
		Value: item,
	}

	if this.tail == nil {
		this.head = &node
		this.tail = &node
		return
	}

	node.Prev = this.tail
	this.tail.Next = &node
	this.tail = &node
}

func (this *DoublyLinkedList[T]) Remove(item T) *T {
	curr := this.head
	for i := 0; curr != nil && i < this.Length; i++ {
		if curr.Value == item {
			break
		}
		curr = curr.Next
	}

	if curr == nil {
		return nil
	}

	return this.removeNode(curr)

}

func (this *DoublyLinkedList[T]) Get(idx int) *T {
	node := this.getAt(idx)
	return &node.Value
}

func (this *DoublyLinkedList[T]) RemoveAt(idx int) *T {
	node := this.getAt(idx)
	if node == nil {
		return nil
	}

	return this.removeNode(node)
}

func (this *DoublyLinkedList[T]) removeNode(node *DllNode[T]) *T {
	this.Length -= 1
	if this.Length == 0 {
		out := this.head.Value
		this.head = this.tail
		this.tail = nil
		return &out
	}

	if node.Prev != nil {
		node.Prev = node.Next
	}

	if node.Next != nil {
		node.Next = node.Prev
	}

	if node == this.head {
		this.head = node.Next
	}
	if node == this.tail {
		this.tail = node.Prev
	}

	node.Prev = node.Next
	node.Next = nil

	return &node.Value
}

func (this *DoublyLinkedList[T]) getAt(idx int) *DllNode[T] {
	curr := this.head

	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.Next
	}

	return curr
}
