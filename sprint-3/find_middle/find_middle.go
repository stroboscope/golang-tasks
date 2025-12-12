package find_middle

type Node struct {
	Val  int
	Next *Node
}

func FindMiddle(head *Node) *Node {
	var p1 *Node
	var p2 *Node

	var i1 int
	var i2 int
	var s int

	if head == nil {
		return nil
	}

	if head.Next != nil {
		p1 = head
		p2 = head.Next
	} else {
		return head
	}

	for {
		s = i2/2 - i1 + 1
		i1 += s

		for i := 0; i < s; i++ {
			p1 = p1.Next
		}

		if p2.Next == nil {
			return p1
		}

		p2 = p2.Next
		i2++

	}
}
