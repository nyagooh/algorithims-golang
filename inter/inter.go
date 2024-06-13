/*
Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

The display will be followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$
*/

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("not enough arguments passed")
	}
	shorter := args[0]
	longer := args[1]
	seen := make(map[rune]bool)
	var result string
	for _, ch := range longer {
		seen[ch] = true
	}
	for _, val := range shorter {
		if seen[val] {
			result += string(val)
			delete(seen, val)
		}
	}
	fmt.Println(result)
	v := map[string]int{
		"nyagooh": 888,
		"bruh":    444,
		"wambo":   666,
		"mom":     777,
	}
	output := SortMap(v)

}

//sorting key and values in a map
func Sortkeys(s map[string]int) {
	keys := make([]string, 0, len(s))
	for i := range s {
		keys = append(keys, i)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, s[k])
	}

}
func SortValue(s map[string]int) {
	values := make([]int, 0, len(s))
	for _, ch := range s {
		values = append(values, ch)
	}
	sort.Ints(values)
	for _, k := range values {
		fmt.Println(k)
	}
}

//merging maps
func MergingMaps(s, b map[string]int) {
	for k, v := range b {
		s[k] = v
	}
	fmt.Println(s)
}
