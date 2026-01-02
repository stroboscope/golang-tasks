package mergelists

type Node struct {
	Val  int
	Next *Node
}

func MergeLists(list1 *Node, list2 *Node) *Node {

	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var Head *Node // голова итогового списка
	l1 := list1    // текущий указатель list1
	l2 := list2    // текущий указатель list2
	l := &Node{}   // пустая нода

	if l1.Val <= l2.Val {
		Head = l1
	} else {
		Head = l2
	}

	if l1.Next == nil && l2.Next == nil {
		if l1.Val <= l2.Val {
			l1.Next = l2
		} else {
			l2.Next = l1
		}
		return Head
	}

	for !(l1 == nil || l2 == nil) {

		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}

	if l1 != nil {
		l.Next = l1
	} else if l2 != nil {
		l.Next = l2
	}

	return Head
}
