package notes

import (
	"fmt"
	"time"
)

func GoChannelNotes() {

}

// 判断channel是否已经关闭
func TestChannelIsClose() {

	c := make(chan int, 3)

	go func() {
		c <- 1
		c <- 2
		close(c)
	}()

	// 让goroutine可以执行完毕
	time.Sleep(time.Second * 1)

	v, close := <-c

	fmt.Println(v)
	fmt.Println(close)
	// 1
	// true

	v, close = <-c

	fmt.Println(v)
	fmt.Println(close)
	// 2
	// true

	v, close = <-c

	fmt.Println(v)
	fmt.Println(close)
	// 0
	// false
}
