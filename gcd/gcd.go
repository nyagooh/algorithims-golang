/*
Write a program that takes two string representing two strictly positive integers that fit in an int.

The program displays their greatest common divisor followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

All arguments tested will be positive int values.

Usage
$ go run . 42 10 | cat -e
2$
$ go run . 42 12
6
$ go run . 14 77
7
$ go run . 17 3
1
$ go run .
$ go run . 50 12 4
$
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	int1, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	int2, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}
	output := Gcd(int1, int2)
	fmt.Println(output)

}
func Gcd(n, v int) int {
	gcd := 1
	for i := 1; i <= n && i <= v; i++ {
		if n%i == 0 && v%i == 0 {
			gcd = i
		}
	}
	return gcd
}
