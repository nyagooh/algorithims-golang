/*Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

If the the argument is not a number, the program should print 00000000.

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args1:= os.Args[1:]
	args := args1[0]
	if len(args1) != 1{
		return
	}
	var new int
	var result string
	// atoi function
	for _, ch := range args {
		if ch < '0' || ch > '9' {
			return
		}
		new = new*10 + int(ch-'0')
	}
	//itoa binary function
	for new > 0 {
		if new % 2 == 0 {
			result = "0"+ result
		}else {
			result = "1" + result
		}
		new /=2
	}
	padder := 8 - len(result)
	for i := 0; i < padder; i++ {
		result = "0" + result
	}
	fmt.Println(result)
}
