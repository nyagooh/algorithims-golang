package main

import (
	"os"
	"strconv"
)

func main() {
	factors := Fprime(8333325)
	result := strconv.Itoa(factors[0])
	for i := 1; i < len(factors); i++ {
		result += "*" + strconv.Itoa(factors[i])
	}
	os.Stdout.WriteString(result + "\n")
}

func Fprime(a int) []int {
	var result []int
	for i := 0; i <= a; i++ {
		if IsPrime(i) && a%i == 0 {
			result = append(result, i)
			a /= i
			if a%i == 0 {
				result = append(result, i)
				a /= i
			}
		}
	}
	return result
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num < 2 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
