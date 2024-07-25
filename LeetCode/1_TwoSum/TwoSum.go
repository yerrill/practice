package twosum

func TwoSum(nums []int, target int) []int {
	valueIndexMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentValue := nums[i]
		requiredValue := target - currentValue

		if foundIndex, ok := valueIndexMap[requiredValue]; ok {
			return []int{i, foundIndex}
		} else {
			valueIndexMap[currentValue] = i
		}
	}

	return nil
}
