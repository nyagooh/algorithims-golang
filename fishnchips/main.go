// Write a function called FishAndChips() that takes an int and returns a string.

//     If the number is divisible by 2, print fish.
//     If the number is divisible by 3, print chips.
//     If the number is divisible by 2 and 3, print fish and chips.
//     If the number is negative return error: number is negative.
//     If the number is non divisible by 2 or 3 return error: non divisible.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}

func FishAndChips(n int) string {
	if n < 1 {
		return "number is negative"
	} else if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	} else if n%2 == 0 {
		return "fish"
	} else if n%3 == 0 {
		return "chips"
	}
	return "non divisible"
}
