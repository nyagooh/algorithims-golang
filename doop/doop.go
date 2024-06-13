// // Write a program that is called doop.

// The program has to be used with three arguments:

// A value
// An operator, one of : +, -, /, *, %
// Another value

// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.
// Usage

// $ go run .
// $ go run . 1 + 1 | cat -e
// 2
// $
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0
// $
// $ go run . 1 % 0 | cat -e
// No modulo by 0
// $
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

// Notions

// Numeric Types
// Arithmetic Operators

	package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	seen := map[string]bool{
		"*": true,
		"/": true,
		"+": true,
		"-": true,
		"%": true,
	}
	if !seen[args[1]] {
		return
	}
	for i, ch := range args[0] {
		if i == 0 && (ch == '+' || ch == '-') {
			continue
		} else if ch < '0' || ch > '9' {
			return
		}
	}
	for i, ch := range args[2] {
		if i == 0 && (ch == '+' || ch == '-') {
			continue
		} else if ch < '0' || ch > '9' {
			return
		}
	}
	num1 := Atoi(args[0])
	// num2 := Atoi(args[2])
	os.Stdout.WriteString(Itoa(num1) + "\n")
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
