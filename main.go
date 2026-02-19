package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	filename := "./big.txt"

	file, err := os.Open(filename)

	log.SetFlags(0)

	if err != nil {
		log.Fatal("failed to read file:", err)
	}

	PrintFileContents(file)
}

func PrintFileContents(file *os.File) {
	const bufferSize = 4096
	//find sectorsize

	buffer := make([]byte, bufferSize)

	for {
		size, err := file.Read(buffer)

		if err != nil {
			break
		}

		fmt.Print(string(buffer[:size]))
	}
}

func CountWords(data []byte) int {
	//counting words
	words := bytes.Fields(data)
	return len(words)
}
