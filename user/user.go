package main

import ("fmt"
"os"
)

func main() {
	User()
}
func User() {
	var name string
	var age int
	fmt.Println("Enter your name:")
	fmt.Scanln( &name)
	fmt.Print("Enter your age:")
	fmt.Scanln( &age)
	if age < 0 || age>100 {
		fmt.Println("Invalid age!")
        os.Exit(1)
	}
	fmt.Printf("Hello %s,your age is %d if i got that right.\n",name,age)
}
