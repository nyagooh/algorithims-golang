package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 5, 9, 7, 5, 7,7}
	fmt.Println("Mode:", FindMode(data))
}

func FindMode(data []int) int {
	occurrence := make(map[int]int)
	for _, num := range data {
		occurrence[num]++
	}
	max := 0
	number := 0
	for num, count := range occurrence {
		if count > max {
			max = count
			number = num
		}
	}
	return number
}
