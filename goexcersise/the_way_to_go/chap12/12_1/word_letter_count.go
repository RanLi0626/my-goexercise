package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('S')
	if err != nil {
		fmt.Printf("got error when read string from input, err is %s", err)
		return
	}
	var runeCount int
	var lineCount int
	wordCount := strings.Fields(str)
	for _, v := range str {
		if v == '\r' || v == '\n' {
			lineCount++
			continue
		}
		runeCount++
	}
	fmt.Printf("runeCount %v, lineCount %v, wordCount %v", runeCount, lineCount, len(wordCount))
}
