// Write a function that takes a string as an argument and returns the letter G if the argument length is less than 3, otherwise returns Invalid Input followed by a newline \n.
//  If it's an empty string return G followed by a newline \n.

package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

func PrintIfNot(str string) string {
	result := "G"
	if len(str)< 1 {
		return result + "\n"
	} else if len(str) < 3 {
		return result + "\n"
	}
	return "invalid input" + "\n"
}