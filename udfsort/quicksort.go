package udfsort

func QuickSort(nums []int) []int {

	numsLen := len(nums)

	start, end := 0, numsLen-1

	quickSortHandle(nums, start, end)

	return nums
}

func quickSortHandle(nums []int, start int, end int) {
	if start >= end {
		return
	}

	middle := partition(nums, start, end)
	quickSortHandle(nums, start, middle-1)
	quickSortHandle(nums, middle+1, end)
}

func partition(nums []int, start int, end int) int {

	pivot := nums[end]

	middle := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[middle], nums[i] = nums[i], nums[middle]
			middle++
		}
	}

	nums[middle], nums[end] = pivot, nums[middle]

	return middle
}
