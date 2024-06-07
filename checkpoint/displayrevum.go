/*Write a program that displays the alphabet in reverse, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

*/
package checkpoint
import "github.com/01-edu/z01"

func ReverseAlp(){
	for i := 'z';i >= 'a';i--{
		if i %2 == 0 {
			z01.PrintRune(i-32)
	} else {
		z01.PrintRune(i)
	}
}
z01.PrintRune('\n')
}