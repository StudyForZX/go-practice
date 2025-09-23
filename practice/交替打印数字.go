package practice

import (
	"fmt"
	"sync"
)

func PrintInAlternationByFor() {
	n := 26

	for i := range n {

		fmt.Print(string('a' + rune(i)))

		fmt.Print(i + 1)
	}
}

func PrintInAlternationBySync() {
	letterCN := make(chan bool)
	numberCN := make(chan bool)

	// 创建一个wati group等待所有goroutine结束
	var wg sync.WaitGroup

	// 启动打印字母的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for letter := 'A'; letter <= 'Z'; letter++ {
			<-letterCN
			fmt.Print(string(letter))
			numberCN <- true
		}

		close(letterCN)
	}()

	// 启动打印数字的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numberCN
			fmt.Print(i)

			if i != 26 {
				letterCN <- true
			}
		}
		close(numberCN)
	}()

	// 交替打印
	letterCN <- true

	wg.Wait()
}
