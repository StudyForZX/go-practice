package ltmath

func LT11MaxAreaByMyself(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	maxArea := 0

	for left < right {

		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		area := min(leftMax, rightMax) * (right - left)

		maxArea = max(maxArea, area)

		if leftMax < rightMax {
			left++
		} else {
			right--
		}

	}

	return maxArea
}
