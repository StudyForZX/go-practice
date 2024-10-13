package udfsort

// 时间复杂度 平均 O(n²) 最好O(n) 最坏O(n²)
// 空间复杂度 O(1)
// 稳定
func BubbleSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums)-1-i; j++ {

			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}
	}

	return nums
}
