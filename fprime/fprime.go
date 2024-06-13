/*Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

    Factors must be displayed in ascending order and separated by *.

    If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

Usage

$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$*/
package main
import (
	// "fmt"
	"os"
	"strconv"
)
func main(){
	factors := Fprime(8333325)
	result := strconv.Itoa(factors[0])
	for i:=1;i<len(factors);i++ {
		result += "*" + strconv.Itoa(factors[i]) 
	}
	os.Stdout.WriteString(result + "\n")
}
func Fprime(n int)[]int{
	result := []int{}
	for i:= 0;i<=n;i++{
		if ISPrime(i) && n%i == 0 {
			result = append(result,i)
			n/=i
		}
		if n%i == 0 {
			result = append(result ,i)
			n/=1
		}
	}
	return result
}
func ISPrime(n int)bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n %3 == 0 || n %2 == 0 {
		return false
	}
	for i := 5;i*i<= n;i+=6 {
		if n%i == 0 && n % (i+2) == 0{
			return false
		}
		
	}
	return true
}





