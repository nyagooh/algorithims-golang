// searchreplace
// Instructions

// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

//     If the number of arguments is different from 3, the program displays nothing.

//     If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage

// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
// $
package main

import (
	"fmt"
	"os"
)

func main(){
	args1 := []byte(os.Args[1:][0])
	args2 := []byte(os.Args[1:][1])
	args3 := []byte(os.Args[1:][2])
	if len(os.Args[1:]) != 3 {
		return
	}
	    for ch := range args2 {
		    for val := range args3 {
				for i:= 0;i < len(args1);i++{
					if args1[i] == args2[ch] {
						args1[i] = args3[val]
					} 
			
				}
			}
		}
		fmt.Println(string(args1))
	}


