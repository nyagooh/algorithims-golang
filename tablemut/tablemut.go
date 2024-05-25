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
	"log"
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
	// var result int
	new := 0
	// atoi function
	for _, ch := range args1 {
		if ch < '0' || ch > '9' {
			log.Fatal("invalid number")
		}
		new = new*10 + int(ch-'0')
	}
	for i := 1; i <= 9; i++ {
		result := new * i
		// output := result + '0'
		fmt.Printf("%d*%d=%d\n",i,new,result)
		// z01.PrintRune(rune(output))
	}
	
// how to range a number when using printrune
	// for _, ch := range fmt.Sprintf("%d",result) {
	// 	z01.PrintRune(ch)
	// }
	
}
