// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to
// chunk a slice into many sub slices where each sub slice has the length of size.
// If the size is 0 it should print a newline ('\n').

package main

import "fmt"


func main() {
		Chunk([]int{}, 10)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	}
func Chunk(slice []int,size int) {
var  result [][]int
 buffer := []int{}
 if size == 0 {
	fmt.Println(slice)
	return
 }
 if len(slice) == 0 {
	fmt.Println(slice)
	return
 }
 for i:= 0;i <= len(slice)-1;i++ {
	if i != 0 && i % size == 0 {
		result = append(result, buffer)
		buffer = []int{}
		buffer = append(buffer, slice[i])
	} else {
		buffer = append(buffer, slice[i])
	}
 }
 if len(buffer) > 0 {
	result = append(result,buffer)
 }
fmt.Println(result)

}
