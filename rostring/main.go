// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.
// Usage

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
// $
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	if len(args) == 0 {
		return
	}
	splits := split(args)
	word :=""
	result :=""
	for i, split := range splits {
		if i == 0 {
			word=split
		}
	}
	for i:= 1; i < len(splits); i++ {
		result +=splits[i] + " "
	}
	fmt.Println(result+word)
}
func split(s string) []string {
	var result []string
	word := ""
	for _, ch := range s {
		if ch != ' ' {
			word += string(ch)
		}
		if ch == ' ' && word != ""{
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	fmt.Println(result)
	return result
}
