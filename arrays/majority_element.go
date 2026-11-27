package arrays

func majorityElement(nums []int) int {
	count := 1
	candidate := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			candidate = nums[i]
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
