package main

import "fmt"

func main(){
	x := 6
	a := &x 				// a is ponting to x's memory address,lets print it out
	fmt.Println(a)
	b := *a 				// asteric means read through the value of this memory address
	fmt.Println(b)
	fmt.Println(*a*b*x) 	// ans will be 6*6*6=216
}
