package main

import "fmt"

func main() {
	fmt.Println("Small number is",min(5,2))
}

func min(num1,num2 int) int { // func-name-parameter with type-return type //
	if (num1 < num2) {
		return num1
	} else {
		return num2
	}
}
