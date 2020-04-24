package main

import "fmt"

func main(){
	//must use make function to create a map//
	Grades := make(map[string]float32)		//map[key_type]return_type//

	Grades["Jim"] = 82
	Grades["Arif"] = 75
	Grades["Joty"] = 90

	fmt.Println(Grades)

	delete(Grades,"Arif")

	fmt.Println(Grades)
}
