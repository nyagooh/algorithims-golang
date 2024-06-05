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
	args1 := []rune(args[0])
	for i,ch := range args1{
			if ch>= 'a' && ch<= 'z' {
				index := 'a' + 'z' -ch
				args1[i] = index
			}else if ch>= 'A' && ch<= 'Z' {
				index := 'A' + 'Z' - ch
				args1[i] = index
			}
		}
		results := string(args1)
	for _, ch := range results {
		z01.PrintRune(ch)
	}
z01.PrintRune(10)
}
