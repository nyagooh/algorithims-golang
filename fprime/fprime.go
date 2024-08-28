/*Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

    Factors must be displayed in ascending order and separated by *.

    If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

Usage

$ 
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
	"fmt"
	"os"
	"strconv"
)

func main() {
args := os.Args[1:]
if len(args) > 1 {
	return
}
factors ,err := strconv.Atoi(args[0])
if err != nil {
	return
}
if factors <= 1 {
    return
}
// var result string
primes := fPrime(factors)
str := itoa(primes[0])
for i:= 1; i<len(primes);i++{
	str+= "*"+itoa(primes[i])
}
fmt.Println(str)
}
func fPrime(n int)[]int{
	var result []int
	for i:=2;i<= n;i++{
		if n%i==0&&IsPrime(i){
			result = append(result,i)
			n/=i
		}
		if n%i == 0{
			result = append(result,i)
			n/=i
		}
		}
	return result
}
func IsPrime(num int)bool{
	
	if num <2{
		return false
	}
	for i:=2;i<num;i++{
		if num%i==0{
			return false
		}
	}
	return true
}
func itoa (num int) string {
	sign := 1
	if num <0{
		sign =-1
		num = num*-1
	}
	var str string
	for num >0 {
		digit := num%10
		str = string(digit+'0') + str
		num/=10
	}
	if sign < 0 {
		str = "-"+str
	}
	return str
}



