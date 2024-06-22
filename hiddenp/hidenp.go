/*
Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

If s1 is an empty string, it is considered hidden in any string.

If the number of arguments is different from 2, the program displays nothing.
Usage

$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
$ go run .
$
*/

package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	args1 := []rune(args[0])
	args2 := []rune(args[1])
	var result []rune
	seen := make(map[rune]bool)
	match := 0
	for _, ch := range args1 {
		seen[ch] = true
		for i := match; i < len(args2); i++ {
			if seen[args2[i]] {
				match = i + 1
				result = append(result, args2[i])
				delete(seen, args2[i])
			}
		}
	}
	if len(result) == len(args1) {
		os.Stdout.WriteString("1")
	} else {
		os.Stdout.WriteString("0")
	}
	os.Stdout.WriteString("\n")
}
