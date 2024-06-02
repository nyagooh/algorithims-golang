package main

import (
	"fmt"
	"os"
)

func main() {
	// found := false
	args := os.Args
	args1 := args[1]
	var result string
	vowels := "aeiouAEIOU"
	for i, ch := range args1 {
		for _, val := range vowels {
			if ch == val{
				results := args1[i:]
				append := args1[:i]
				result = results + append + "ay"
			}
			
		}
		
	}
	if len(result) != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("no vowels")
	}
	
}
