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
	"fmt"
	"log"

	// "math/big"
	"os"
	"strconv"
)

func main() {
	ar := os.Args[1:]
	if len(ar) == 0 {
		log.Fatal("invalid number of arguments")
		return
	}
	args := ar[0]
	args1 := ar[1]
	args3 := ar[2]
	if len(args) == 0 {
		log.Fatal("invalid number of arguments")
		return
	}
	num, err := strconv.Atoi(args)
	if err != nil {
		log.Fatalf("string convert atoi has an error")
	}
	num1, err := strconv.Atoi(args3)
	if err != nil {
		log.Fatalf("string convert atoi has an error")
	}
	punctuation := []string{"+", "-", "*", "/", "%"}
	for _, k := range punctuation {
		if args1 != k {
			return
		}
	}
	if num < 0 && num > 9 {
		return
	}
	if num1 < 0 && num1 > 9 {
		return
	}
	var result int
	maxint := 9223372036854775807
	if num == maxint && num1 != 0 && args1 != "*" {
		return
	}
	switch args1 {
	case "+":
		result = num + num1
		fmt.Printf("%d", result)
		fmt.Println()
	case "-":
		result = num - num1
		fmt.Printf("%d", result)
		fmt.Println()
	case "*":
		result = num * num1
		fmt.Printf("%d", result)
		fmt.Println()
	case "/":
		if num1 == 0 {
			fmt.Println("No division by 0")
			// return
		}
		result = num / num1
		fmt.Printf("%d", result)
		fmt.Println()
	case "%":
		if num1 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = num % num1
		fmt.Printf("%d", result)
		fmt.Println()
	}
	fmt.Println()

}
