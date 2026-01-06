package main

import (
	"bytes"
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
	//counting words
	words := bytes.Fields(data)
	return len(words)
}
