package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Word Counter Begins...")
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	//specific edge case - putting guards on top of the function
	if len(data) == 0 {
		return 0
	}

	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}

	return wordCount
}
