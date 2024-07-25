package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		out    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, v := range cases {
		result := TwoSum(v.in, v.target)

		if (result[0] == v.out[0] && result[1] == v.out[1]) || (result[1] == v.out[0] && result[0] == v.out[1]) {
			t.Logf("PASS TwoSum(%v, %v) == %v, got %v", v.in, v.target, v.out, result)
		} else {
			t.Errorf("FAIL TwoSum(%v, %v) == %v, got %v", v.in, v.target, v.out, result)
		}

	}
}
