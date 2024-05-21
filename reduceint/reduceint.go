<<<<<<< HEAD
// the function should have as paremeters a slice a[]int and a function f  func(int,int)int.For each element of the slice ,it should apply the function
// f func(int,int)int.save the result and then print it
package main

import (
	// "fmt"

	"fmt"

	"github.com/01-edu/z01"
)
=======
// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice,
// it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

package main

import "fmt"
>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd

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
<<<<<<< HEAD
	as := []int{500, 2}
=======
	as := []int{500, 2,}
>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
<<<<<<< HEAD
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
=======



func ReduceInt(a []int, f func(int, int) int) {
  char := a[0]
  var result int
  for i:= 0; i < len(a);i++ {
	result = f(char,a[i])

  }
fmt.Println(result)
}

>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd
