// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

// The case of the letter remains unchanged, for example :

// 'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

// The final result will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program prints a new line.
// Usage

// $ go run . "abc"
// zyx$
// $
// $ go run . "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// $
// $ go run .
// $
package main

import (
	// "fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune(10)
	}
	// array1 := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	// array2 := []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	// arraycap := []rune{}
	 for _, ch := range args {
		r := []rune(ch)
		for i,val := range r {
			if val >= 'a'&& val <= 'z' {
				index := 'a' + 'z' - val
				r[i] = index
			} else if val >= 'A' && val <= 'Z' {
				opposite := 'A' + 'Z' - val
				r[i] = opposite
			}

		}
		result := string(r)
	    for _, ch := range result {
			z01.PrintRune(ch)
		}
	 }
	 z01.PrintRune(10)

}
