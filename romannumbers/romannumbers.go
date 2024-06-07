package main

import "fmt"

func main() {
 number := 78
 fmt.Println(intToRoman(number))
}

func intToRoman(num int) string{
	// if num <= 0 || num >= 4000 {
	// 	return "", fmt.Errorf("ERROR: cannot convert to roman digit")
	// }

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanNum := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			romanNum += romans[i] + "+"
			num -= vals[i]
		}
	}
	romanNum = romanNum[:len(romanNum)-1]
	return romanNum
}
