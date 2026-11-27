package arrays

func removeDuplicates(nums []int) int {
	i := 0
	for _, v := range nums {
		if nums[i] != v {
			i++
			nums[i] = v
		}
	}
	return i + 1
}
