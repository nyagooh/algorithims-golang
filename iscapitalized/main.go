// Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter
//  or a non-alphabetic character.

//     If any of the words begin with a lowercase letter return false.
//     If the string is empty return false.

// Expected function

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
func IsCapitalized(s string) bool {
var check  bool
word := ""
var result[]string
for _,v := range s {
	if string(v) != " " {
		word+=string(v)
	}else {
		result = append(result, word)
		word = ""
	}
}
if word != ""{
	result = append(result, word)
}
for _, ch := range result {
	if ch[0] >= 'a' && ch[0] <= 'z' {
		check=false
	}else {
		check=true
	}
}
return check
}
