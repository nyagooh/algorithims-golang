//write a function that returns last rune of a string

package checkpoint


func LastRune(str string)rune{
	m := []rune(str)
	return m[len(m)-1]

}