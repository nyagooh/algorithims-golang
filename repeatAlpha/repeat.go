/*Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

The result must be followed by a newline ('\n').

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

If the number of arguments is different from 1, the program displays nothing.

Usage
$ go run . abc | cat -e
abbccc
$ go run . Choumi. | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
$ go run . "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
$ go run .
$ go run . "" | cat -e
$
$
*/

package main

import "os"

func main(){
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	args1 := args[0]
	if len(args1) < 1 {
		return
	}
	output := repeat_alpha(args1)
	os.Stdout.WriteString(output)
	os.Stdout.WriteString("\n")
}
func repeat_alpha(s string)string{
	var count int
	var result string
	for _,ch := range s {
		if ch >= 'a' || ch <= 'z' {
			count = int(ch-'a') + 1
			for i:= 1;i<=count;i++{
				result +=string(ch)
			}
		} else  if ch >= 'A' || ch <= 'Z' {
			count = int(ch-'A') + 1
			for i:= 1;i<=count;i++{
				result +=string(ch)
			}
		}
		result+=string(ch)
	}
	return result
}