package main

import (
	"fmt"
	"linked_list"
)

func printLL(ll *linked_list.LinkedList) {
	runner := ll.Head
	for runner != nil {
		fmt.Print(runner.Val)
		fmt.Print(" -> ")
		runner = runner.Next
	}
	fmt.Println("nil")
}

func main() {
	testList := &linked_list.LinkedList{}
	testList.Head = &linked_list.ListNode{1, nil}
	testList.Head.Next = &linked_list.ListNode{2, nil}

	fmt.Println(testList.Contains(2))

	testList.Insert(3)

	fmt.Println(testList.Contains(3))

	printLL(testList)

	testList.AddFront(0)

	printLL(testList)

	testList.Insert(4)

	printLL(testList)
}
