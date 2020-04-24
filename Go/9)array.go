package main

import "fmt"

func main() {
	var lis1 = []int{1,2,3,45,5}	//initialization of array with only elements//
	var lis2 = [5]int{5,1,2,3,6}	//initialization of array with size and elements//
	var lis3 [5]int 				//declaration of array with size//

	for i := 0; i < 5; i++ {
		lis2[i] = lis2[i]+lis1[i]
	}
	fmt.Println(lis2)
	lis3 = lis2
	fmt.Println(lis3)
}
