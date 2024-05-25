// chunk
// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to
// chunk a slice into many sub slices where each sub slice has the length of size.
// If the size is 0 it should print a newline ('\n').
// Expected function
// func Chunk(slice []int, size int) {
// }
// Usage
// Here is a possible program to test your function :
// package main
// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
// And its output :
// $ go run .
// []
// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
// $ }

package main

import (
	"fmt"

	// "strings"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)

	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {

	if size == 0 {
		fmt.Println(slice)
		return
	}

	var result [][]int
	 for i := 0 ; i <= len(slice)-1;i+=size {
		end := i+ size
		if end  > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	 }
	 fmt.Println(result)

	// buffer := []int{}
	// for i := 0; i < len(slice); i++ {
	// 	if i != 0 && i%size == 0 {
	// 		result = append(result, buffer)
	// 		buffer = []int{}
	// 		buffer = append(buffer, slice[i])
	// 	} else {
	// 		buffer = append(buffer, slice[i])
	// 	}
	// }

	// if len(buffer) > 0 {
	// 	result = append(result, buffer)
	// }
	//  print2DArray(result)
}
// func print2DArray(a [][]int) {
// 	z01.PrintRune('[')
// 	for i, innerSlice := range a {
// 		if i > 0 {
// 			z01.PrintRune(' ')
// 		}
// 		z01.PrintRune('[')
// 		for j, num := range innerSlice {
// 			if j > 0 {
// 				z01.PrintRune(' ')
// 			}
// 			z01.PrintRune(rune(num) + '0')
// 		}
// 		z01.PrintRune(']')
// 	}
// 	z01.PrintRune(']')
// 	z01.PrintRune(10)
// }