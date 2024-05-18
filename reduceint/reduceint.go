// the function should have as paremeters a slice a[]int and a function f  func(int,int)int.For each element of the slice ,it should apply the function
// f func(int,int)int.save the result and then print it
package main

import (
	// "fmt"

	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) > 0 {
		z01.PrintRune(' ')
	}
	input := a[0]
	 for i := 1;i < len(a);i++ {
		input = f(input, a[i])
	 }
	 fmt.Println(input)
}
