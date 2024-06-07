// the function should have as paremeters a slice a[]int and a function f  func(int,int)int.For each element of the slice ,it should apply the function
// f func(int,int)int.save the result and then print it

// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice,
// it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

package main

import "github.com/01-edu/z01"

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
	var result int
	for i := 1; i < len(a); i++ {
		result = f(input, a[i])
	}
	//itoa
	output := itoa(result)
	for _, ch := range output {
		z01.PrintRune(ch)
	}

}
func itoa(s int) string {
	sign := 1
	str := ""
	if s < 1 {
		s *= -1
		sign *= -1
	}
	for s > 0 {
		digit := s % 10
		str = string(digit + '0') + str
		s /= 10
	}
	if sign < 0 {
		str = "-" + str
	}
	return str
}
// func atoi(s string) int {

// 	str := 0
// 	for _, ch := range s {
// 		if ch < '0' || ch > '9' {

// 		}
// 		str = str*10 + int(ch-'0')
// 	}
// 	return str
// }
