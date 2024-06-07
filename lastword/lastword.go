// Write a program that takes a string and displays its last word, followed by a newline ('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// Usage

package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
	// "strings"
	// "github.com/01-edu/z01"
	// "strings"
	// "github.com/01-edu/z01"
	// "github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	str := ""
	for i := len(args) - 1; i >= 0; i-- {
		if string(args[i]) == " " && str == "" {
			continue
		} 
		 if string(args[i]) != " "  {
			str = string(args[i])
		}
		if string(args[i]) == " " {
			break
		}
		


	}
	
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)
}
