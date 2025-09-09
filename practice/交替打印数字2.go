package practice

import (
	"fmt"
)

func PrintNumAndChar() {

	numChan := make(chan bool)
	charChan := make(chan bool)
	doneChan := make(chan bool)

	go printNum(numChan, charChan)
	go printChar(charChan, numChan, doneChan)

	numChan <- true

	<-doneChan

	close(numChan)
	close(charChan)
	close(doneChan)
}

func printNum(numChan chan bool, charChan chan bool) {
	for i := range 10 {
		<-numChan
		fmt.Printf("Num is %d\n", i)
		charChan <- true
	}
}

func printChar(charChan chan bool, numChan chan bool, doneChan chan bool) {
	for i := range 10 {
		<-charChan
		fmt.Printf("Char is %c\n", rune('a'+i))
		if i != 9 {
			numChan <- true
		}
	}

	doneChan <- true
}
