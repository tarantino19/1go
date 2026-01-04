package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Word Counter Begins...")
	data, _ := os.ReadFile("./words.txt")

	fmt.Println("data:", string(data))

}
