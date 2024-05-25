// foldint
// Instructions

// The function should have as parameters a function, f func(int, int) int a slice of integers, slice []int and an int acc int.
// For each element of the slice, it should apply the arithmetic function, save the result and print it.
// The function will be tested with our own functions Add, Sub, and Mul.
// The function should have as parameters a function, f func(int, int) int a slice of integers, slice []int and an int acc int. For each element of the slice, it should apply the arithmetic function, save the result and print it. The function will be tested with our own functions Add, Sub, and Mul.
// Expected function

// func FoldInt(f func(int, int) int, a []int, n int) {

// }

// Usage

// Here is a possible program to test your function:

// package main

// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }

// And its output :

// $ go run .
// 99
// 558
// 87

// 93
// 0
// 93
// $
package main

import (
	"fmt"
	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
func FoldInt(f func(int, int) int, a []int, n int) {
	var result int
	var new int
	for _, ch := range a {
		new += ch
	}
	result = f(new, n)

	if result < 0 {
		result = result * -1
	}
	results := Itoa(result)
	for _, ch := range results {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)	

}
func Add(a, b int) int {
	return a + b
}
func Mul(a, b int) int {
	return a * b
}
func Sub(a, b int) int {
	return a - b
}
func Itoa(a int) string {
	sign := 1
	if a == 0 {
		return "0"
	}
	if a < 0 {
		sign *= -1
		a *= -1
	}
	str := ""
	for  a > 0 {
		digit := a % 10
		str = string(digit+'0') + str
		a /= 10
	}
	if sign == -1 {
		str = "-" + str
	}
	return str
}
