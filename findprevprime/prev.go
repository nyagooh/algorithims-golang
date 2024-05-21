// findprevprime
// Instructions

// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.
// Expected function

// func FindPrevPrime(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "fmt"

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }

// And its output :

// $ go run .
// 5
// 3
// $
package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	if nb > 1 && nb%2 != 0 || nb == 2 {
		return nb
	} else if nb > 1 && nb%2 == 0 && nb != 2 {
		return nb - 1
	}
	return nb
}
