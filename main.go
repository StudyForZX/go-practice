package main

import (
	"fmt"

	"github.com/studyforzx/lt/ltmath"
)

func main() {

	// res := ltmath.LengthOfLongestSubstring("xajlkhjghxkkhlahk3")
	res := ltmath.MinWindows("ADOBECODEBANC", "ABC")

	fmt.Println(res)

}
