/*
Write a function that takes a byte, swaps its halfs (like the example) and returns the result.

Expected function
func SwapBits(octet byte) byte {

}
*/
package main

import "fmt"

func main(){
	var octet byte  =  0x14
	swapped := SwapBits(octet)
	fmt.Printf("%x/n",swapped)
}
func SwapBits(octet byte)byte{
	return octet <<4|octet >>4
}