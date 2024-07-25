package addtwonumbers

import "testing"

func ArrayToListNode(array []int) *ListNode {
	var output *ListNode

	for i := len(array) - 1; i >= 0; i-- {
		if output == nil {
			output = &ListNode{array[i], nil}
		} else {
			output = &ListNode{array[i], output}
		}
	}

	return output
}

func ListNodeToArray(list *ListNode) []int {
	output := []int{}
	l := list

	for l != nil {
		output = append(output, l.Val)
		l = l.Next
	}

	return output
}

func CompareArrays(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}

	return true
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1     []int
		l2     []int
		result []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, v := range cases {
		l1 := ArrayToListNode(v.l1)
		l2 := ArrayToListNode(v.l2)
		result := addTwoNumbers(l1, l2)
		resultArray := ListNodeToArray(result)

		if CompareArrays(v.result, resultArray) {
			t.Logf("PASS addTwoNumbers(%v, %v) == %v, got %v", v.l1, v.l2, v.result, resultArray)
		} else {
			t.Errorf("FAIL addTwoNumbers(%v, %v) == %v, got %v", v.l1, v.l2, v.result, resultArray)
		}
	}
}
