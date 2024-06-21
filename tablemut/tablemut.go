// tabmult
// Instructions

// Write a program that displays a number's multiplication table.

//     The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.

// Usage

// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $
// loop over this
package main

import (
	"fmt"
	// "log"
	"os"

	// "github.com/01-edu/z01"
	// "github.com/01-edu/z01"
)



func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided. Usage: go run . <argument>")
		return
	}
	args1 := args[0]
	if len(args1) == 0 {
		return
	}
	var output string
	for i := 1; i <= 9; i++ {
		result := Atoi(args1) * i
		str := Itoa(result)
		str2 := Itoa(i)
		output = str2 + " x "+ args1 + " = " + str
		os.Stdout.WriteString(output)
		os.Stdout.WriteString("\n")
		
	}
}
func Atoi(n string) int {
	sign := 1
	var result int
	if n[0] == '-' {
		sign *= -1
		n = n[1:]
	}
	for _, ch := range n {
		result = result*10 + int(ch-'0')
	}
	return result * sign
}

func Itoa(n int) string {
	var result string
	sign := 1
	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign *= -1
		n *= -1
	}
	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n /= 10
	}
	if sign < 0 {
		result = "-" + result
	}
	return result
}

