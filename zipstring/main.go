package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	var result string
		for i := 0; i < len(s); i++ {
		count := 1                           // Start count at 1
		for i+1 < len(s) && s[i] == s[i+1] { // Check for consecutive occurrences
			count++
			i++
		}
		result += strconv.Itoa(count) + string(s[i])
		 // Move to the next character (after the counted occurrences)
	}
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))                                   // Output: 1Y2o3u1n1g1F3e3l1a1s
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog")) // Output: 1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
	fmt.Println(ZipString("Helloo Therre!"))                                    // Output: 1H1e2l2o1 1T1h1e2r1!
}