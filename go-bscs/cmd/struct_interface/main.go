package main

import "fmt"

type user struct {
	names string
	car
}

type gasEngine struct {
	mpg uint8
	gallons uint8
}

type evEngine struct {
	mkwh uint8
	kwh uint8
}

type car struct {
	brand_name string
	engine gasEngine
}

func (g gasEngine) milesLeft() uint8{
	return g.gallons*g.mpg
}

func (ev evEngine) milesLeft() uint8{
	return ev.kwh*ev.mkwh
}


type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles<=e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("you can't make it there!")
	}
}




func main() {
	//var car1 car = car{brand_name: "Mercedes Benz" }
	var user1 user = user{"Serge Danson", car{brand_name:"fdfd"} }

	fmt.Printf("user details: %v", user1.names);
	//fmt.Printf("car details: %v", car1);

	//unknonimous struct == not reusable

	var myEngine2 = struct{
		mpg uint8
		galoon uint8
	} {2, 5}
	fmt.Println(myEngine2)

	var myEnginer3 evEngine = evEngine{mkwh: 3, kwh: 20}

	println(myEngine2.galoon)

	canMakeIt(myEnginer3, 40)
}