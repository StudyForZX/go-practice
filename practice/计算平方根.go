package practice

func practiceSquareRoot(x int) int {

	left := 0
	right := x
	middle := 0

	for left <= right {

		middle = (left + right) / 2

		if middle*middle > x {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}

	return right
}
