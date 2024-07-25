// Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

//     The numbers must be separated by a comma and a space.
//     If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
//     Prepend a 0 to any number that is less than 10.
//     Add a new line \n at the end of the string.

// Expected function

package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	var result string
	if from >99 || from < 1 || to >99 ||to <1{
		return "Invalid" + "\n"
	}
	for from < to {
		str := itoa(from)
		if len(str)<2 {
			str="0"+str
		}
		result+=str+", "
		from++
	}
	for from > to {
		str := itoa(from)
		if len(str)<2 {
			str="0"+str
		}
		result+=str+", "
		from--
	}
	if from==to{
		str := itoa(from)
		if len(str)<2 {
			str="0"+str
		}
		result+=str+"\n"
	}
	return result
}

func itoa(s int) string {
	var result string
	for s > 0 {
		digit := s % 10
		result = string(digit+'0') + result
		s/= 10
	}
	return result
}
