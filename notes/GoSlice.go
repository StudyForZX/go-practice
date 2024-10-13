package notes

import "fmt"

// 切片的数据结构
// type SliceHeader struct {
// 	Data uintptr // 指向数组的指针
// 	Len  int	 // 当前切片的长度
// 	Cap  int	 // 当前切片的容量
// }

func GoSliceNotes() {

	// go切片笔记
	// https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/
	// 特点：动态数组、长度不固定、容量不足会自动扩容

	// 声明：
	// 切片类型的声明方式与数组有一些相似，不过由于切片的长度是动态的，
	// 所以声明时只需要指定切片中的元素类型
	s1 := []int{}
	s2 := []interface{}{}

	fmt.Println(s1, s2)

	// 初始化
	s3 := []int{1, 2, 3}
	s4 := make([]int, 10)
	fmt.Println(s3, s4)

	// 访问元素
	s3Len := len(s3)
	s3Cap := cap(s3)
	fmt.Println(s3Len, s3Cap)

	// 切片的追加或者扩容机制
	s3 = append(s3, 4)
	fmt.Println(s3)
	// 扩容原理
	// 1. 如果期望容量大于当前容量的两倍就会使用期望容量。
	// 2. 如果当前切片的长度小于1024就会将容量翻倍。
	// 3. 如果当前的切片长度大于1024就会每次增加25%的容量，知道新容量大于预期容量。

	// 切片拷贝
	copy([]int{}, s3)
}
