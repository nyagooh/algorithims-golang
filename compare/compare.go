// compare
// Instructions

// Write a function that behaves like the Compare function.
// Expected function

// func Compare(a, b string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Compare("Hello!", "Hello!"))
// 	fmt.Println(piscine.Compare("Salut!", "lut!"))
// 	fmt.Println(piscine.Compare("Ola!", "Ol"))
// }

// And its output :

// $ go run .
// 0
// -1
// 1
// $

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
func Compare(a, b string) int {
	// count:= 0
	// for char := range a{
	// 	for i:=0;i<len(b);i++{
	// 		count++
	// 		if b[i] == a[char]&& len(a) == len(b){
	// 			return 0
	// 		} else if  b[i] != a[char]&& len(a) != len(b) {
	// 			return -1
	// 		} else if   b[i] == a[char]&& len(a) != len(b) {
	// 			return 1
	// 		}
	// 	}
	// }
	// return count
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}
