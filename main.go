package main

import (
	ds "Playground/data-structures"
	"fmt"
)

type TestStruct struct {
    Name string
}

func main() {
    fmt.Println("Linked List")
    LinkedList()
    fmt.Println("")

//    fmt.Println("Doubly Linked List")
//    DoublyLinkedList()
}

func DoublyLinkedList(){
    dll := ds.DoublyLinkedList[TestStruct]{}

    dan :=  TestStruct {
        Name: "Daniel",
    }
    dll.Append(dan)

    peter :=  TestStruct {
        Name: "Peter",
    }
    dll.Append(peter)

    ant :=  TestStruct {
        Name: "Anthony",
    }
    dll.Append(ant)

    x := dll.Get(0)
    fmt.Println(x)
    dll.RemoveAt(2)

    y := dll.Get(2)
    fmt.Println(y)
}

func LinkedList(){
    ll := ds.LinkedList[TestStruct]{}

    dan :=  TestStruct {
        Name: "Daniel",
    }
    ll.Add(dan)

    peter :=  TestStruct {
        Name: "Peter",
    }
    ll.Add(peter)

    ant :=  TestStruct {
        Name: "Anthony",
    }
    ll.Add(ant)

    ll.Remove(dan)

    fmt.Println(ll)
}

