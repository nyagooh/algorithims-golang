package checkpoint

import (
	"github.com/01-edu/z01"
)

func Hello(){
	v := "hello world"
	for _, ch := range v {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}