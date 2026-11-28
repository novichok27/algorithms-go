package arrays

func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	for idx, val := range nums {
		tmp[(idx+k%n)%n] = val
	}
	copy(nums, tmp)
}
