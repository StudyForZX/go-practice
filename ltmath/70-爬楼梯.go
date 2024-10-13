package ltmath

// 递归
func LT70_ClimbStairsByRecursion(n int) int {

	res := 0

	for i := 1; i <= n; i++ {
		res = climbStair(n)
	}

	return res
}

func climbStair(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStair(n-1) + climbStair(n-2)
}

// 循环
func LT70_ClimbStairsByFor(n int) int {

	x := 0
	y := 1

	for i := 0; i < n; i++ {

		x, y = x+y, x

	}

	return x
}
