// write a function that takes a slice of int and an int to remove from the slice, and returns the slice without the int.
//usage DeleteInt([]{1,2,3,4,5}, 3) expected []{1,2,4,5}

package main

import "fmt"

func main(){
	fmt.Println(DeleteInt([]int{1,2,3,4,5}, 3))
}

func DeleteInt(slice []int, i int) []int {
	var result []int
	for k := 0; k < len(slice); k++ {
		if slice[k] == i {
			result = append(result,slice[:k]...)
			result = append(result, slice[k+1:]...)

		} 
	}
 return result
}

func Add(a, b int) int {
	return a + b
}