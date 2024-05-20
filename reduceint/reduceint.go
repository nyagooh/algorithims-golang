// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice,
// it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

package main

import "fmt"

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
	as := []int{500, 2,}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}



func ReduceInt(a []int, f func(int, int) int) {
  char := a[0]
  var result int
  for i:= 0; i < len(a);i++ {
	result = f(char,a[i])

  }
fmt.Println(result)
}

