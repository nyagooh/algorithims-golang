package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output"
	}
	result := ""
	store := split(str)
	//fmt.Println(store)
	for i:=len(store)-1; i >= 0; i--{
		result = result + store[i] + " " 
		

	}
	result = result + "\n"
	return result
}

func split(s string) []string{
	word := ""
	var result []string
	for _, value := range s{
		if value == ' '{
			if word != ""{
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(value)
		}
	}
	if word != ""{
		result = append(result, word)
	}
	return result
}
