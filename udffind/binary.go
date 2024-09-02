package udffind

// 二分法 返回boolean类型
func BinarySearchReturnBoolean(arr []int, goal int) bool {

	arrLen := len(arr)

	if arrLen == 0 {
		return false
	}

	middle := arrLen / 2

	if arr[middle] == goal {
		return true
	}

	left := arr[0:middle]
	right := arr[middle+1:]

	resLeft := BinarySearchReturnBoolean(left, goal)
	resRight := BinarySearchReturnBoolean(right, goal)

	return resLeft || resRight
}

// 二分法 返回索引
func BinarySearchReturnIndex(arr []int, goal int) int {

	arrKeys := []int{}

	for k := range arr {
		arrKeys = append(arrKeys, k)
	}

	return binarySearchReturnIndexHandle(arr, arrKeys, goal)
}

func binarySearchReturnIndexHandle(arr []int, arrKeys []int, goal int) int {

	arrLen := len(arrKeys)

	if arrLen == 0 {
		return -1
	}

	middle := arrLen / 2

	if (arr[arrKeys[middle]]) == goal {
		return arrKeys[middle]
	}

	left := arrKeys[0:middle]
	right := arrKeys[middle+1:]

	leftRes := binarySearchReturnIndexHandle(arr, left, goal)
	rightRes := binarySearchReturnIndexHandle(arr, right, goal)

	if leftRes >= 0 {
		return leftRes
	}

	if rightRes >= 0 {
		return rightRes
	}

	return -1
}
