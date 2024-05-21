package main

import "fmt"

func main() {
	fmt.Println(prevPrime(5))
	fmt.Println(prevPrime(4))
	fmt.Println(prevPrime(2))
}
func prevPrime(nb int) int {
	if nb % 2 != 0 && nb > 0 || nb == 2{
		return nb
	} else {
		return nb-1
	}
}