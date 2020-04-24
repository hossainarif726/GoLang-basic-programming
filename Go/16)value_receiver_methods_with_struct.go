package main

import "fmt"

const u16bitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {gas_pedal uint16
				break_pedal uint16
				steering_wheel int16
				top_speed_kmh float64}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax/kmh_multiple)
}

func main(){
	first_car := car{gas_pedal : 65535,
					break_pedal : 0,
					steering_wheel :13566,
					top_speed_kmh : 225.0}
	fmt.Println(first_car.kmh())
	fmt.Println(first_car.mph())
}
