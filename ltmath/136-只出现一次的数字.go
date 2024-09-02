package ltmath

func SingleNumber(nums []int) int {

	tmp := 0

	for _, num := range nums {
		tmp = tmp ^ num
	}

	return tmp
}
