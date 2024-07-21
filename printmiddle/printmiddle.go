// Write a program that prints the middle argument of the command line.

// If the number of arguments is less than 1, print ('\n')
// If the number of arguments is even, print the middle two arguments with a space (' ') between the arguments
// If the number of arguments is odd, print the middle one.
// Print ('\n') at the end of the output.

package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main(){
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune(10)
		return
	}
	var result string
	if len(args)% 2 == 0 {
		new := len(args)/2
		result = args[new-1] + " " + args[new] + "\n"
	} else {
		new := len(args) / 2
		result = args[new] + "\n"
	}
	os.Stdout.WriteString(result)
}
