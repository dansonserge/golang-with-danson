package main

import "fmt"

func square(item2 *[5]float32) [5]float32{
	for i := range item2 {
		item2[i] = item2[i] * item2[i]
	}

	return *item2
}

func main() {
	var item1 = [5]float32{23,34,23,54,42}
	fmt.Printf("Value of item1 before: %v", item1)

	var item2 = square(&item1)
	fmt.Printf("value of item1 after: %v", item2)
	fmt.Printf("value of item2: %v", item2)

}