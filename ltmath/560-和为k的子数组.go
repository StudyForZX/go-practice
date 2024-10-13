package ltmath

func LT560_SubarraySumByPreSum(nums []int, k int) int {

	sum := 0
	count := 0
	preSumMap := map[int]int{}
	preSumMap[0] = 1

	for _, num := range nums {

		sum += num

		if preSumMap[sum-k] > 0 {
			count += preSumMap[sum-k]
		}

		preSumMap[sum]++
	}

	return count
}

func LT560_SubarraySumByFor(nums []int, k int) int {

	count := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}
