package main

import (
	"fmt"
	"os"
)


func main(){
	args := os.Args
	shorter := args[1]
	longer := args[2]
	seen := make(map[rune]bool)
	var result string
	for _, ch := range longer {
		seen[ch] = true
	}
	for _,val := range shorter {
		if seen[val] {
			result+=string(val)
			delete(seen,val)
		}
	}
	fmt.Println(result)
}