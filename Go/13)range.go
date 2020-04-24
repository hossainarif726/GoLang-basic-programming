package main

import "fmt"

func main() {
	num := []int{1,3,4}

	for i:= range num {
		fmt.Println("Elements","[",i,"]","is",num[i])
	}
}
