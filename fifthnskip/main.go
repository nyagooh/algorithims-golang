package main

import "fmt"

func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
}
func FifthAndSkip(str string) string {
	if len(str) == 0 || len(str) == 5 {
		return ""
	}
	words := split(str)
	var result string
	for i, char := range words {
		if i%5 == 0 && i != 0 {
			result += " "
		}
		if i != 4 {
			result += string(char)
		}
	}
	return result
}
func split(str string) []string {
	var result []string
	word := ""
	for _, value := range str {
		if value != ' ' {
			word += string(value)
		}
		if value == ' ' {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
//