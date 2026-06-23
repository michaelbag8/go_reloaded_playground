package main

import (
	"fmt"
	"os"
)

//Mini Task 1: Read and Write Files
func main() {
	if len(os.Args) != 3 {
		fmt.Println("\033[1;31mUsage: go run sample.txt output.txt\033[0m")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err!=nil{
		fmt.Printf("\033[31m%v\033[0m\n", err)
		return
	}

	err = os.WriteFile(output, []byte(data), 0644)
	if err!=nil{
		fmt.Printf("\033[31m%v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32mWritten to output file successfully\033[0m")
}
