// Write a program that takes a string and displays its last word, followed by a newline ('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// Usage

package main

import (
	// "fmt"
	"os"
	// "strings"

	"github.com/01-edu/z01"
)
func main(){
	args := os.Args[1:]
	args2 := args[0]
	var count int = 1
	var index int
	for i,ch := range args2 {
		if ch == ' ' {
			count++
			index = i
		}
	}
   result := (args2[index:]) 
   for _,ch := range result {
	z01.PrintRune(ch)
   }
}
