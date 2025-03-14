package main

import "fmt"


type gasEngine struct {
	galoons float32
	mpg float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine]struct {
	carMake string
	carModel string
	engine T
}

func structGenetics() {
	var gasCar =  car[gasEngine]{
		carMake: "Rolls Royce",
		carModel: "Cullinan",
		engine: gasEngine{
			galoons: 12.4,
			mpg: 48.0,
		},
	}

	var evCar =  car[electricEngine]{
		carMake: "tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh: 12.4,
			mpkwh: 48.0,
		},
	}

	fmt.Printf("car1: %v", gasCar.carMake)
	fmt.Printf("car2: %v", evCar.carMake)
}


