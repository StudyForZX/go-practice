package ltmath

import (
	"math"
	"strings"
)

func LT8MyAtoi(str string) int {
	var handleSign bool
	var sign bool = true
	var numStrBuilder strings.Builder
	var zeroStart bool = true

	for _, s := range str {

		if !handleSign && s == ' ' {
			continue
		}

		if !handleSign {

			handleSign = true

			if s == '-' {
				sign = false
				continue
			} else if s == '+' {
				sign = true
				continue
			} else if s < '0' || s > '9' {
				return 0
			}
		}

		if s >= '0' && s <= '9' {

			if zeroStart && s == '0' {
				continue
			} else {
				zeroStart = false
			}

			numStrBuilder.WriteRune(s)

		} else {
			break
		}
	}

	res := 0
	numStr := numStrBuilder.String()
	for _, nChar := range numStr {

		d := int(nChar - '0')

		if sign {
			if res > (math.MaxInt32-d)/10 {
				return math.MaxInt32
			}
			res = res*10 + d
		} else {
			if res < (math.MinInt32+d)/10 {
				return math.MinInt32
			}
			res = res*10 - d
		}
	}

	return res
}
