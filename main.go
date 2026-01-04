package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Word Counter Begins...")
	data, _ := os.ReadFile("./words.txt")

	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}

	wordCount++
	fmt.Println(wordCount)
}
