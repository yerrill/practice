package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var output *ListNode
	var lastNode *ListNode
	carry := 0
	var l1val, l2val, value int

	for l1 != nil || l2 != nil {

		// Check nil values and advance nodes
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		} else {
			l1val = 0
		}

		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		} else {
			l2val = 0
		}

		// Calculate Value and carry

		value = (l1val + l2val + carry) % 10
		carry = (l1val + l2val + carry) / 10

		// Append Node

		node := &ListNode{value, nil}

		if output == nil {
			output = node
			lastNode = node
		} else {
			lastNode.Next = node
			lastNode = node
		}
	}

	// Additional node if carry remaining

	if carry > 0 {
		lastNode.Next = &ListNode{carry, nil}
	}

	return output
}
