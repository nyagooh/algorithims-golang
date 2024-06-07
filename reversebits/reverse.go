/*Write a function that takes a byte, reverses it bit by bit (as shown in the example) and returns the result.

Expected function
func ReverseBits(oct byte) byte {

}
*/

package main

import "fmt"

func main() {
	// Example byte to reverse
	var original byte = 0b00101100 // You can change this value to test other cases
	reversed := ReverseBits(original)
	fmt.Printf("Original: %08b\n", reversed)
}

func ReverseBits(oct byte) byte {
	var reversed byte
	for i := 0; i < 8; i++ {
	// reversed |= (oct & (1 << i)) >> i << (7-i)
		reversed = (reversed << 1) | (oct & 1) 
		oct >>= 1
	}
	return reversed
}
