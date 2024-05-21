package main

import (
	"fmt"
	"os"
)

func main() {
	// array1 := []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','s','t','u','v','w','x','y','z'}
	// array2 := []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	args := os.Args[1:][0]
	if len(args) != 1 {
		fmt.Println("")
	}
    arraycap := []rune{}
	var result string
	for _, ch := range args {
		if ch >= 'a' && ch <= 'z' {
			index := 'a' + 'z' - ch
			 arraycap = append(arraycap, index)
		}
		result = string(arraycap)
	}
	fmt.Println(result)
	// for _, ch := range args {
	// 	r := []rune(ch)
	// 	for i,val := range r {
	// 		if val >= 'a'&& val <= 'z' {
	// 			index := 'a' + 'z' - val
	// 			r[i] = index
	// 		} else if val >= 'A' && val <= 'Z' {
	// 			opposite := 'A' + 'Z' - val
	// 			r[i] = opposite
	// 		}

	// 	}
	// 	result := string(r)
	//     for _, ch := range result {
	// 		z01.PrintRune(ch)
	// 	}
	//  }
	//  z01.PrintRune(10)

}



