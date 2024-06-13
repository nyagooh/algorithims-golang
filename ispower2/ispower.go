/*
Write a program that determines if a given number is a power of 2. A number n is a power of 2 if it meets the following condition: n = 2 m where m is another integer number.

For example, 4 = 2 2 or 16 = 2 4 are power of 2 numbers while 24 is not.

This program must print true if the given number is a power of 2, otherwise it has to print false.

If there is more than one or no argument, the program should print nothing.

When there is only one argument, it will always be a positive valid int.
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
$
*/

package main

import (
	"fmt"
	"log"
	"os"
)
//bitwise operation
func main() {
	// atoi function
	var new int
	for _, ch := range os.Args[1] {
		if ch < '0' || ch > '9' {
			log.Fatal("not an integer")
		}
		new = new*10 + int(ch-'0')
	}

	fmt.Println(Ispower(new))
	fmt.Println(iIspower(new,2))
}

func Ispower(n int) bool {
	if n > 0 && n&(n-1) == 0 {
		return true
	}
	return false
}
func iIspower (n , base int)bool{
	if n == 1 {
		return true
	}
	for n % base == 0 {
		n /= base 
		if n == 1 {
			return true
		}
	}
	return false
}
