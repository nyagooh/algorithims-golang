package main
import "fmt"
func main() {
	j := 9
	s := []int{2,3,4,5}
	fmt.Print(SliceInt(s,j))
}
func SliceInt(s []int,j int)[]int{
	var index []int
	for x,num := range s {
		for i:=0;i<len(s);i++{
			if num + s[i] ==j && x != i{
				index = append(index,x,i)
				return index
			} else{
				continue
			}
		} 
	}
	return nil
}