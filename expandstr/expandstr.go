// Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

// The string will be followed by a newline ('\n').

// A word, in this exercise, is a sequence of visible characters.

// If the number of arguments is not 1, or if there are no word, the program displays nothing.

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	if len(args) < 1 {
		return
	}
	// var result[]string
	var result2 []string
	word := ""
	for _, ch := range args {
		if string(ch) == " " {
			if word != "" {
				result2 = append(result2, word)
				word = ""
			}

		} else {
			word += string(ch)
		}
	}
	if word != "" {
		result2 = append(result2, word)
	}
	fmt.Println(result2)
	result := ""
	for i := 0; i < len(result2)-1; i++ {
		result += result2[i] + "   "
		

	}
	 result+= result2[len(result2)-1]
	fmt.Println(result)
}
