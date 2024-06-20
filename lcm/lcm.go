package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	int1, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	int2, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}
	output := Lcm(int1,int2)
	fmt.Println(output)
}

func Lcm (n,v int)int {
gcd := 1
var lcm int
	for i :=1; i <= v && i <= n; i++{
		if n % i == 0 && v % i == 0 {
			gcd = i
		}
		lcm = (v*n) /gcd
	}
	return lcm
}