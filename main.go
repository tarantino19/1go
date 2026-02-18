package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Word Counter Program Begins...")

	data, err := os.ReadFile("./words1.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	//counting words
	words := bytes.Fields(data)
	return len(words)
}
