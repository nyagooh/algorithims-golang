package main

import (
	"fmt"
	"os"
)


func main(){
	args := os.Args
	args1 := args[1]
	args2 := args[2]
	seen := make(map[rune]bool)
	var filter string

	for _, ch := range args1 + args2 {
		if !seen[ch]{
			seen[ch] = true
			filter += string(ch)
		}
	}
	fmt.Println(filter)
}