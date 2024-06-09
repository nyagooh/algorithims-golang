// write A function that Behaves like the compare function
// expected output go run . 0 -1 1
package main

import "fmt"

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut", "lut"))
	fmt.Println(Compare("Ola", "Ol"))
}

func Compare(a, b string) int {
	A := []rune(a)
	B := []rune(b)
	for i := 0; i < len(A) && i < len(B); i++ {
		if A[i] != B[i] {
			if A[i] < B[i] {
				return -1
			} else if A[i] > B[i] {
				return 1
			}
		}
	}
	if len(A) > len(B) {
		return 1
	}
	return 0
}
