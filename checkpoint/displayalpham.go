package checkpoint

import "github.com/01-edu/z01"

func DisplayAlpham(){
	for i:= 'a'; i <= 'z';i++ {
		if i%2 == 0 {
			z01.PrintRune(i-32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}