package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Println("Number =",a)  //for loop//
	}

	//In golang there is no while loop//
	//Thats why we are goinf to use for loop as while loop//
	b := 25
	for true {
		fmt.Println(b)
		b--				// while loop implementation //
		if (b == 0) {
			break
		}
	}
}
