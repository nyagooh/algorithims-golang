<<<<<<< HEAD
package main

import ( "os"
"github.com/01-edu/z01"
)


func main(){
	args := os.Args[1:][0]
	for _,char := range args {
		if char == ' ' {
			 char = char + 1
			z01.PrintRune(' ')

		} 
		if char == ',' {
			z01.PrintRune(',')
		}
		if char >= 'a'&& char <= 'z' {
			newchar := (char-'a'+13)%26 + 'a'
			z01.PrintRune(newchar)
		} else if char >= 'A' && char <= 'Z' {
			new := (char-'A' + 13)%26 + 'A'
			z01.PrintRune(new)
		}
	}
	z01.PrintRune(10)
=======
// Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

//     'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

//     The output will be followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays nothing.

// Usage

// $ go run . "abc"
// nop
// $ go run . "hello there"
// uryyb gurer
// $ go run . "HELLO, HELP"
// URYYB, URYC
// $ go run .
package main

import ("os"
"github.com/01-edu/z01"
)

func main(){
args := os.Args[1:]
if len(args) != 1 {
return
}
args2 := args[0]
for _, ch := range args2 {
	if ch == ' ' {
		z01.PrintRune(' ')
	}
	if ch ==  ','{
		z01.PrintRune(',')
	}
	if ch >= 'a' && ch <= 'z' {
		result := (ch- 'a'+13)%26 + 'a'
		z01.PrintRune(result)
	} else if ch >= 'A' && ch <= 'Z' {
		output := (ch-'A'+13)%26 + 'A'
		z01.PrintRune(output)
	}
}
z01.PrintRune(10)
>>>>>>> 20834b4cdd843a04453382b23f41f0633496c6dd
}