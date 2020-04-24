package main

import "fmt"

//structure(struct) is a user defined type in golang//
//which can take various type unlike array//

type car struct {gas_pedal uint16
				break_pedal uint16
				steering_wheel int16
				top_speed float64}

func main(){
	first_car := car{gas_pedal : 32456,
					break_pedal : 0,
					steering_wheel :3566, //using our own type car//
					top_speed : 226.52}

	second_car := car{gas_pedal : 35436,
					break_pedal : 3,
					steering_wheel : 3766,
					top_speed : 521.52}

	fmt.Println(first_car.steering_wheel)
	fmt.Println(second_car.top_speed)

}
