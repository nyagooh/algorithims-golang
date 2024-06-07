// /Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing./ 
// "faya" "fghaufghysdera"

package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	arg1 := os.Args[1]
// 	arg2 := os.Args[2]

// 	str := ""

// 	for i, v := range arg2 {
		
// 	}

// }

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		// z01.PrintRune()
		return
	}

	args1 := []rune(os.Args[1:][0])
	args2 := []rune(os.Args[1:][1])
	match := 0
	// lens := 0
	var result []rune
	for _, char := range args1 {
		for i := match; i <= len(args2)-1; i++ {
			if char == args2[i] {
				match = i+1
				result = append(result, char)
				break
			}
		}
	}
	
	// if lens == len(args1) {
	// 	for _, ch := range args1 {
	// 		z01.PrintRune(ch)
	// 	}
	if len(result) == len(args1) {
		for _, ch := range result {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
	
}
