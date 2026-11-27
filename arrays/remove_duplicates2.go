package arrays

func removeDuplicates2(nums []int) int {
	k := 0
	for _, v := range nums {
		if k < 2 || nums[k-2] != v {
			nums[k] = v
			k++
		}
	}
	return k
}
