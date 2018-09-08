package linked_list

type ListNode struct {
	Val int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Insert(val int) {
	if ll.Head == nil {
		ll.Head = &ListNode{val, nil}
	}

	runner := ll.Head

	for runner.Next != nil {
		runner = runner.Next
	}

	runner.Next = &ListNode{val, nil}
}

func (ll *LinkedList) Contains(target int) bool {
	if ll.Head == nil {
		return false
	}

	runner := ll.Head

	for runner != nil {
		if runner.Val == target {
			return true
		}

		runner = runner.Next
	}

	return false
}

func (ll *LinkedList) AddFront(target int) {
	if ll.Head == nil {
		ll.Head = &ListNode{target, nil}
		return
	}

	new_node := &ListNode{target, ll.Head}
	ll.Head = new_node
}

func (ll *LinkedList) Remove(target int) {
	runner := ll.Head
	prev := &ListNode{}

	for runner.Next != nil {
		if runner.Val == target {
			prev.Next = runner.Next
		}
		prev = runner
		runner = runner.Next
	}
}





