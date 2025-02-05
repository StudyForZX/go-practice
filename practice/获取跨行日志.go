package practice

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getFileContent() *os.File {
	const FILE_PATH = "/Users/bytedance/Desktop/mutil-line.log"
	f, err := os.Open(FILE_PATH)
	if err != nil {
		panic("Error reading file:" + err.Error())
	}
	return f
}
func GetMutiLineLog() {

	fileContent := getFileContent()
	// 将日志按行分隔
	scanner := bufio.NewScanner(fileContent)
	logPattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \[\w+\] .*`)
	currentLog := ""

	for scanner.Scan() {
		line := scanner.Text()
		if logPattern.MatchString(line) {
			if currentLog != "" {
				fmt.Println("Parsed Log Entry:\n", currentLog)
				currentLog = ""
			}
			currentLog += line + "\n"
		} else {
			currentLog += line + "\n"
		}
	}

	if currentLog != "" {
		fmt.Println("Parsed Log Entry:\n", currentLog)
	}
}
