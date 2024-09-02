package ltmath

func Convert(s string, numRows int) string {

	if numRows <= 1 {
		return s
	}

	sArray := []byte(s)
	resArray := make([][]byte, numRows)
	down := true
	index := -1
	res := ""

	for _, word := range sArray {

		if down {
			index++
		} else {
			index--
		}

		if index == numRows-1 {
			down = false
		} else if index == 0 {
			down = true
		}

		resArray[index] = append(resArray[index], word)

	}

	for _, item := range resArray {
		res += string(item)
	}

	return res
}
