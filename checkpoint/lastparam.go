package checkpoint

import (
	"os"

	"github.com/01-edu/z01"
)

func lastParam() {
	args := os.Args[1:]
	if len(args) > 1 {
		for _, char := range args {
			for _, value := range char {
				if char  == args[len(args)-1] {
					z01.PrintRune(value)
				} 
			}
		}
	} 
	z01.PrintRune('0')
	z01.PrintRune('\n')
}
