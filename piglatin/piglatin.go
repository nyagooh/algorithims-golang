// Instructions
// Write a program that transforms a string passed as argument in its Pig Latin version.

// The rules used by Pig Latin are as follows:

// If a word begins with a vowel, just add "ay" to the end.
// If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
// Only the latin vowels will be considered as vowel(aeiou).
// If the word has no vowels, the program should print "No vowels".

// If the number of arguments is different from one, the program prints nothing.

// Usage
// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	args1 := args[0]
	vowels := "aeiouAEIOU"
	str := ""
	foundVowel := false

	for i := range args1 {
		for _, ch := range vowels {
			if string(args1[0]) == string(ch) {
				str = args1 + "ay"
				foundVowel = true
				break

			} 
			 if  string(args1[i]) == string(ch) {
				str = args1[i:] + args1[:i] + "ay"
				foundVowel = true
				break
			} 
			
		}
		if foundVowel {
			break
		}
	}
	if !foundVowel {
		str = "no vowels found"
	}
	fmt.Println(str)
}
