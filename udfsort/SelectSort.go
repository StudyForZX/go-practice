package udfsort

// 时间复杂度 平均 O(n²) 最好O(n²) 最坏O(n²)
// 空间复杂度 O(1)
// 不稳定
func SelectSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}

		}

	}

	return nums
}

func SelectSort2(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {

		minIndex := i

		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
