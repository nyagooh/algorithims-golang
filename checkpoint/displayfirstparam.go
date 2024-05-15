package checkpoint

import ("os"
        "github.com/01-edu/z01"
)

func FirstParam() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, ch := range args[0] {
			z01.PrintRune(ch)
		}
	} else {
		z01.PrintRune('0')
	}
z01.PrintRune('\n')
}