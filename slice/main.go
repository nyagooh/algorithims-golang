package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 1 {
		nbrs = append(nbrs, len(a))
	}
	if nbrs[0] > nbrs[1]{
		return nil
	}
	if nbrs[0] < 0{
		nbrs[0] = len(a) + nbrs[0]
	}
	if nbrs[1] < 0{
		nbrs[1] = len(a) + nbrs[1]
	}
	var result []string
	
		result = append(result, a[nbrs[0]:nbrs[1]]...)
	return result
}