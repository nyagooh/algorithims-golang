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
}