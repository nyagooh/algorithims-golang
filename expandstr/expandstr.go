// Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

// The string will be followed by a newline ('\n').

// A word, in this exercise, is a sequence of visible characters.

// If the number of arguments is not 1, or if there are no word, the program displays nothing.

package main

import (
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
		if string(ch) != " " {
			word += string(ch)
		} else {
			result2 = append(result2, word)
			word = ""
		}
	}
	if word != "" {
		result2 = append(result2, word)
	}
	for i, ch := range result2 {
		if i != len(result2)-1{
		os.Stdout.WriteString(ch + "   ")
		}
		if i == len(result2)-1 {
			os.Stdout.WriteString(ch + "\n")
		}
	}
}
