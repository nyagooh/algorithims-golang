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
	if nb <= 1 {
		return 0
	}
	prime := nb
	for !IsPrime(prime) {
		prime--
	}
	return prime
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num < 2 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
