package ltmath

// 递归
func LT70ClimbStairsByMyselfWithRecursion(n int) int {
	return climbStair(n)
}

func climbStair(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return climbStair(n-1) + climbStair(n-2)
}

// 循环
func LT70ClimbStairsByMyselfWithFor(n int) int {

	x := 0
	y := 1

	for range n {
		x, y = y, x+y
	}

	return x
}
