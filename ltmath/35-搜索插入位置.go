package ltmath

func LT35SearchInsertByMyself(nums []int, target int) int {

	for index, num := range nums {

		if num < target {
			continue
		}

		return index

	}

	return len(nums)

}

func LT35SearchInsertByFor(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		middle := (left + right) / 2

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}
