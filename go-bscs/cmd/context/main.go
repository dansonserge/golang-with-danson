package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	openConnectionWithTimeout()
}


func openConnection(done chan bool){
	fmt.Println("attentempting connection...")

	if rand.IntN(100)>50 {
		fmt.Println("oops! hanging connection!")
		time.Sleep(100000 * time.Hour)
	}else{
		time.Sleep(2 * time.Second)
		fmt.Println("connection Established")
	}
	done <- true
}

func openConnectionWithTimeout(){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer  cancel()

	done := make(chan bool)
	go openConnection(done)

	select {
		case <- done:
			fmt.Println("Connection Successful")
		case <- ctx.Done():
			fmt.Println("connection Timeout!")
	}
}