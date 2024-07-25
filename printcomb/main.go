// Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

// These combinations are separated by a comma and a space.

package main

import "fmt"

func main() {
	PrintComb()
}

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 1; j <= 9; i++ {
			for z := 2; z <= 9; z++ {
				if i < j && j<z {
					fmt.Print(i)
					fmt.Print(j)
					fmt.Print(z)
					if i <= 7 {
						fmt.Print(",")
						fmt.Print( )
					}
				}
			}
		}
	}
	fmt.Println()
}
