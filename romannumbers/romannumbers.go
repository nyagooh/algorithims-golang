package main

import "fmt"

func main() {
 number := 999
 str1, str2 := toRoman(number)
 fmt.Println(str1)
 fmt.Println(str2)
}
	func toRoman(num int) (string, string) {
		var result, breakdown string
	
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		calculations := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	
		for i, value := range values {
			if num >= value {
				num = num -value
				result += symbols[i]
				if len(breakdown) > 0 {
					breakdown += "+"
				}
				breakdown += calculations[i]
			}
		}
		return breakdown,result
	}