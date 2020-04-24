
package main

import "fmt"

func main() {
	var grade string
	marks := 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default : grade = "D"  //if none of the case works default will be triggerd//
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B":
		fmt.Println("Good")
	case grade == "C":
		fmt.Println("Average")
	default :
		fmt.Println("Poor result")
	}
}
