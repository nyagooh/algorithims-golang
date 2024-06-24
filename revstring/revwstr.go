package main

import "fmt"

func main(){
	s := "How are you"
	word := ""
	str := []string{}
	for _, v := range s {
		if v != ' ' {
			word+=string(v)
		} else {
			str = append(str, word)
			word = ""
		}
	}
	if word != "" {
		str = append(str, word)
	}
	result:= ""
	for i:= len(str)-1;i>=0;i--{
		if result != "" {
			result+=" "
		}
		result+=str[i]
	}
	fmt.Println(result)
}