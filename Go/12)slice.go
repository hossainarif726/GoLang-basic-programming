package main

import "fmt"

//slice overcome the limitation of dynamic size of array//

func main() {
	arr := make([]int,2,3) 	//a slice of len 2 and possible number of element accomodation//
	var arr2 []int

	arr = append(arr,1,3,4) 	//appending three numbers//
	fmt.Println(cap(arr)) 	//total capacity of arr slice//
	fmt.Println(arr)
	arr2 = arr[1:]		//index slicing with slice//
	fmt.Println(arr2)
}
