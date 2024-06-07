package main

import (
	"fmt"
	"os"

	// "github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
	}
	result := len(args)-1
	count := 0
	for range args {
	  count++
	if count == result {
	new := args[count]
	fmt.Print(new)
	fmt.Println()
	}
	}
}