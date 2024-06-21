/*Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

If the number of arguments is different from 1, the program displays nothing.
Error cases have to be handled as shown in the example below.
Usage
$ go run . 10
a
$ go run . 255
ff
$ go run . 5156454
4eae66
$ go run .
$ go run . "123 132 1" | cat -e
ERROR$
$
*/

package main

import (
	"fmt"
	"os"
)

func main() {
  args := os.Args[1:]
  if len(args) != 1{
	return
  }
  for _, ch := range args[0] {
	if ch == ' '{
		fmt.Println("ERROR")
		return
	}
  }
  input := atoi(args[0])
  output := DecimalToHex(input)
  fmt.Println(output)

}

func DecimalToHex(n int) string {
	if n == 0 {
		return "0"
	}
	hexString := ""
	hex := "0123456789abcdef"
	for n > 0 {
		reminder := n % 16
		hexString = string(hex[reminder]) + hexString
		n /= 16
	}
	return hexString
}
 func atoi(str string)int {
	var new int
	for _, ch := range str {
		if ch < '0' || ch > '9'{
			return 0
		}
		new = new * 10 + int(ch -'0')
	}
	return new
 }
