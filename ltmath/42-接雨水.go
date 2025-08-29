package ltmath

func LT42_Trap(heights []int) int {

}

func LT42_Trap_ByMySelf(heights []int) int {
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
