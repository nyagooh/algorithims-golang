// Write a program that prints the ASCII value of a letter passed in the command line

// If the argument is not a letter nothing will be printed
// if the number of arguments is not 1 then nothing will be printed

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 1 {
		return
	}
	args1 := args[0]
	var result int
	for _, ch := range args1 {
		if ch < 32 || ch > 126  {
			return
		} else if ch >= 48 || ch <= 57 {
			return
		}
		result = int(ch)
	}
	fmt.Println(result)
}
