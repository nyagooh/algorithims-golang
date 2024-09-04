package main

import "fmt"

func main() {
	s := "hello world"
	sb := "hel"
	fmt.Println(substring(s, sb))
	fmt.Println(split(s))
	fmt.Println(reverse(s))
	fmt.Println(Doubles(s))
	fmt.Println(Repeat(s))
}

// function for substring
func substring(s string, sb string) string {
	lens := len(s)
	lensb := len(sb)
	var str string
	for i := 0; i < lens-lensb; i++ {
		if s[i:i+lensb] == sb{
			str += s[i:i+lensb] 
		}
	}
	return str
}

//function for splitting
func split(s string)[]string{
var result []string
word := ""
for _, value := range s {
	if value != ' '{
		word += string(value)
	}
	if value ==' '{
        result = append(result, word)
        word = ""
    }
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}

//function for finding doubles
func Doubles(value string)bool{
	var double bool
	for i:=0;i<len(value)-1;i++{
		if value[i] == value[i+1] {
			double = true
            break
		}
	}
	return double
}
//checking for repeated values which dont necessarry follow each other
func Repeat(value string)string{
repeat := make(map[rune]int)
for _, v := range value {
		repeat[v]++
}
var str string
maxcount := 0
for value,count := range repeat {
	if count > maxcount {
		maxcount = count
        str = string(value)
	}
}
return str
}
//function for reverse
func reverse(value string) string {
	word := split(value)
	reverse := ""
	for i:=len(word)-1;i>=0;i-- {
		if reverse != ""{
			reverse += " "
		}
		reverse += word[i]
		}
		return reverse
	}

//function for concating// reverse concatenate // concatenate//

