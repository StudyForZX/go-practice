package ltmath

// 递归
// func LT70_ClimbStairs(n int) int {

// 	res := 0

// 	for i := 1; i <= n; i++ {
// 		res = climbStair(n)
// 	}

// 	return res
// }

// func climbStair(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	if n == 2 {
// 		return 2
// 	}

// 	return climbStair(n-1) + climbStair(n-2)
// }

func LT70_ClimbStairs(n int) int {

	x, y, sum := 0, 0, 1

	for i := 1; i <= n; i++ {
		x = y
		y = sum
		sum = x + y
	}

	return sum
}
