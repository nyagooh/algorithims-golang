package checkpoint

import (
	"os"

	"github.com/01-edu/z01"
)

func lastParam() {
	args := os.Args[1:]
	new := args[len(args)-1]
	if len(args) > 1 {
		for _ , ch := range new {
			z01.PrintRune(ch)
		}
	z01.PrintRune('0')
	z01.PrintRune('\n')
}
}
