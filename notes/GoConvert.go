package notes

import (
	"fmt"
	"strconv"
)

func GoConvertNotes() {

	// go的整形转浮点型
	var i1 int = 42
	var f1 float64 = float64(i1)
	fmt.Println(f1)

	// go的浮点型转整形
	var f2 float64 = 3.14159
	var i2 int = int(f2)
	fmt.Println(i2)

	// 字符串转整形
	var s3 = "123"
	i3, err := strconv.Atoi(s3)
	if err == nil {
		fmt.Println(i3)
	}
	// 或者使用parseInt
	i4, err := strconv.ParseInt(s3, 10, 64)
	if err == nil {
		fmt.Println(i4)
	}

	// 整形转字符串
	i5 := 123
	s5 := strconv.Itoa(i5)
	fmt.Println(s5)

	// go不支持直接对bool转int
}
