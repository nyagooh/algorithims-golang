package checkpoint

func StringLength(str string)int{
	// runes := []rune(str)
	// return len(runes)
	count := 0
	for range str{
		count++
	}
	return count
	
}