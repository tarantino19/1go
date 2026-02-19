package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile("./words.txt")

	log.SetFlags(0)

	if err != nil {
		log.Fatal("failed to read file:", err)
	}

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	//counting words
	words := bytes.Fields(data)
	return len(words)
}
