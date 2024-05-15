package main

import (
	v "check/checkpoint"
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := "Tandiwe and nyagooh"
	first := v.FirstRune(str)
	last := v.LastRune(str)
	z01.PrintRune(first)
	z01.PrintRune('\n')
	z01.PrintRune(last)
	z01.PrintRune('\n')
	v.DisplayAlpham()
	s := "Hello world"
	l := v.StringLength(s)
	fmt.Println(l)
	v.ReverseAlp()
	v.FirstParam()
	lastParam()
	v.Hello()
	v.ParamCount()
	a := []int{123,445,667,888,999}
	j := v.Max(a)
	fmt.Println(j)

}
func lastParam() {
	args := os.Args[1:]
	for _, char := range args {
		for _, value := range char {
			if char == args[len(args)-1]{
				z01.PrintRune(value )
			} 
			
		}
	}
	z01.PrintRune('\n')
}
