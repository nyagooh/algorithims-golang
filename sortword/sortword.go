/*
sortwordarr
Instructions
Write a function SortWordArr that sorts by ascii (in ascending order) a string slice.

Expected function
func SortWordArr(a []string) {

}
Usage
*/
package main

import "fmt"

func main() {
	words := []int{2,5,7,8,3,5}
	SortWordArr(words)
	fmt.Println(words) // Output: [apple banana cherry]
}
func SortWordArr(words []int) {
	//initialize the first loop
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if words[j] > words[j+1] {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
}
