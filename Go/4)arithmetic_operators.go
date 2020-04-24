package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a)   //user input//
	fmt.Scan(&b)

	fmt.Println("Additon =",a+b)
	fmt.Println("Subtraction =",a-b)
	fmt.Println("Multiplication =",a*b)
	fmt.Println("Division =",a/b)
	fmt.Println("Modulus =",a%b)
}
