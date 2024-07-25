package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	bestLeft, bestRight, left, right := 0, 1, 0, 1
	charsInSubstring := make(map[byte]int)

	charsInSubstring[s[left]] += 1

	for right < len(s) && left < len(s) {
		count := charsInSubstring[s[right]]

		if count <= 0 {
			charsInSubstring[s[right]] += 1
			right += 1

			if right-left > bestRight-bestLeft {
				bestLeft = left
				bestRight = right
			}
		} else {
			charsInSubstring[s[left]] -= 1
			left += 1
		}
	}

	return bestRight - bestLeft
}
