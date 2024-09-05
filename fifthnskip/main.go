package main

import "fmt"

func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	var sentence string
	var result string
	var count int = 0
	for _, c := range str {
		if c == ' ' {
			continue
		} else {
			sentence += string(c)
		}
	}

	for i, c := range sentence {
		count++
		if i == 0 {
			result += string(c)
			continue
		}
		if count%6 == 0 {
			result += string(" ")
			count = 0
		} else {
			result += string(c)
		}

	}
	return result
}
