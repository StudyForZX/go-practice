package udfsort

// 时间复杂度 平均 O(n²) 最好O(n) 最坏O(n²)
// 空间复杂度 O(1)
// 稳定
func BubbleSort(nums []int) []int {

	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}

	for i := range numsLen {
		for j := i + 1; j < numsLen; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
