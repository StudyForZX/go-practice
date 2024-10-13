package udfsort

// 以切片方式 空间利用率低
func QuickSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	middle := partition(nums)
	QuickSort(nums[:middle])
	QuickSort(nums[middle+1:])

	return nums
}

func partition(nums []int) int {

	// 选择最尾部元素为基准
	pivot := len(nums) - 1
	middle := 0

	for i := 0; i < pivot; i++ {
		// 这里控制排序的顺序
		if nums[i] < nums[pivot] {
			nums[i], nums[middle] = nums[middle], nums[i]
			middle++
		}
	}

	nums[middle], nums[pivot] = nums[pivot], nums[middle]

	return middle
}
