// Write a program that takes a string and displays its last word, followed by a newline ('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// Usage

package main

import (
	// "fmt"
	"os"
	"strings"
	"github.com/01-edu/z01"
	// "strings"
<<<<<<< HEAD

	"github.com/01-edu/z01"
=======
	// "github.com/01-edu/z01"
>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd
)
func main(){
	args := os.Args[1:]
	if len(args) != 1{
		return
	}
	args2 :=strings.TrimSpace(args[0])
	if args2 == "" {
		return
	}
	var count int = 1
	var index int
	for i,ch := range args2 {
		if ch == ' ' {
			count++
			index = i
		}
	}
   result := (args2[index:]) 
<<<<<<< HEAD
   for _,ch := range result {
	z01.PrintRune(ch)
   }
=======
   result2 := strings.TrimSpace(result)
   for _,ch := range result2 {
	z01.PrintRune(ch)
   }
   z01.PrintRune(10)
>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd
}
