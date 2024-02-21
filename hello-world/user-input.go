package main

import "fmt"

func main() {
	var fName string
	fmt.Println("Enter first name:")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("Enter last name")
	fmt.Scanln(&lName)
	fmt.Println("Full name is: " + fName + " " + lName)
}
