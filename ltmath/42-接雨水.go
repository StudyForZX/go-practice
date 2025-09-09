package ltmath

func LT42TrapByDoublePoints(heights []int) int {

	count := 0
	lPoint := 0
	rPoint := len(heights) - 1
	lMax := 0
	rMax := 0

	for lPoint < rPoint {

		lMax = max(lMax, heights[lPoint])
		rMax = max(rMax, heights[rPoint])

		if heights[lPoint] < heights[rPoint] {
			count += lMax - heights[lPoint]
			lPoint++
		} else {
			count += rMax - heights[rPoint]
			rPoint--
		}

	}

	return count
}

func LT42TrapByMonotonicStack(heights []int) int {
	count := 0
	stack := []int{}

	for i, height := range heights {
		for len(stack) > 0 && height > heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(heights[left], height) - heights[top]
			count += curWidth * curHeight
		}
		stack = append(stack, i)
	}

	return count
}

// Dynamic Programming
// 空间复杂度是 O(n)
func LT42TrapByDP(heights []int) int {

	heightsLen := len(heights)
	if heightsLen < 3 {
		return 0
	}

	count := 0

	lMax := make([]int, heightsLen)
	rMax := make([]int, heightsLen)
	lMax[0] = heights[0]
	rMax[heightsLen-1] = heights[heightsLen-1]

	for i := 1; i < heightsLen; i++ {
		lMax[i] = max(lMax[i-1], heights[i-1])
	}

	for i := heightsLen - 2; i > 0; i-- {
		rMax[i] = max(rMax[i+1], heights[i+1])
	}

	for i := range heightsLen {
		current := min(lMax[i], rMax[i]) - heights[i]
		if current > 0 {
			count += current
		}
	}

	return count
}

func LT42TrapByN2(heights []int) int {

	heightsLen := len(heights)
	if heightsLen < 3 {
		return 0
	}

	count := 0

	for i := range heightsLen {
		// 找到左边最高的柱子
		lMax := 0
		for j := range i {
			lMax = max(lMax, heights[j])
		}

		// 找到右边最高的柱子
		rMax := 0
		for j := i + 1; j < heightsLen; j++ {
			rMax = max(rMax, heights[j])
		}

		tmp := min(lMax, rMax) - heights[i]
		if tmp > 0 {
			count += tmp
		}
	}

	return count
}

func LT42TrapByMySelf(heights []int) int {
	var (
		count      int = 0
		lPoint     int = 0
		rPoint     int = 1
		heightsLen int = len(heights)
		tmpLen     int = 0
	)

	for rPoint < heightsLen {

		if heights[lPoint] > heights[rPoint] {
			tmpLen++
		} else {

			var delete = 0
			for i := lPoint + 1; i < rPoint; i++ {
				delete += heights[i]
			}

			count += min(heights[rPoint], heights[lPoint])*tmpLen - delete
			tmpLen = 0
			lPoint = rPoint
		}

		rPoint++
	}

	// 以下为AI补充
	// 处理剩余部分：当右侧没有比左侧更高的柱子时
	// 从右向左重新处理剩余部分
	if lPoint < heightsLen-1 {
		rPoint = heightsLen - 1
		lPoint = rPoint - 1
		tmpLen = 0

		for lPoint >= 0 {
			if heights[rPoint] >= heights[lPoint] {
				tmpLen++
			} else {

				var delete = 0
				for i := lPoint + 1; i < rPoint; i++ {
					delete += heights[i]
				}
				count += min(heights[rPoint], heights[lPoint])*tmpLen - delete
				tmpLen = 0
				rPoint = lPoint

			}

			lPoint--
		}
	}

	return count
}
