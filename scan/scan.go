package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scanf("Enter your name: %s", &name)
	fmt.Scanf("Enter your age: %d", &age)

	fmt.Println("You are:", name, "And you are:", age, "years old!")
}
