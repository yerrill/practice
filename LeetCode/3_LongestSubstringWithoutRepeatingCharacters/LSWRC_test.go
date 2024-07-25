package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLongestSubstring(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, v := range cases {
		result := lengthOfLongestSubstring(v.input)

		if v.output == result {
			t.Logf("PASS lengthOfLongestSubstring(%v) == %v, got %v", v.input, v.output, result)
		} else {
			t.Errorf("FAIL lengthOfLongestSubstring(%v) == %v, got %v", v.input, v.output, result)
		}
	}
}
