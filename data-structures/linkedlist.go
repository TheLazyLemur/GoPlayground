package datastructures

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (this Node[T]) String() string {
    return fmt.Sprintf("%v", this.Value)
}

type LinkedList[T comparable] struct {
	Length int
	head   *Node[T]
}

func (this *LinkedList[T]) Add(value T) {
    newNode := new(Node[T])
    newNode.Value = value

    if this.head == nil {
        this.head = newNode
    }else{
        it := this.head
        for ; it.Next != nil; it = it.Next { }
        it.Next = newNode
    }

    this.Length ++
}

func (this *LinkedList[T]) Remove(value T) {
    var previous *Node[T]

    for current := this.head; current != nil; current = current.Next{
        if current.Value == value {
            if this.head == current {
                this.head = current.Next
            }else{
                previous.Next = current.Next 
                return
            }
        }
        previous = current
    }
}

func (this LinkedList[T]) Get(value T) *Node[T] {
    for it := this.head; it != nil; it = it.Next{
        if it.Value == value {
            return it
        }
    }

    return nil
}

func (this LinkedList[T]) GetAt(idx int) *Node[T] {
    curr := this.head
    for i := 0; i < idx; i++ {
        curr = curr.Next
    }
    return curr
}

func (this LinkedList[T]) String() string {
    sb := strings.Builder{}
    
    for it := this.head; it != nil; it = it.Next{
        sb.WriteString(fmt.Sprintf("%v ", it))
    }

    return sb.String()
}

