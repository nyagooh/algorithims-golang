package main

func main(){
//atoi function
}

func Ispower(number ,base int)bool {
	if number == 1{
		return true
	}
	for number%base == 0 {
		number /= base
		if number == 1{
			return true
		}
	}
	return false
}