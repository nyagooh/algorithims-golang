// write a function that behaves like the compare function
// expected output go run . 0 -1 1
package main

import "fmt"

func main(){
	fmt.Println(Compare("Hello!","Hello!"))
	fmt.Println(Compare("Salut","lut"))
	fmt.Println(Compare("Ola","Ol"))

}
func Compare(a,b string)int{
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}