package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1,3,4}
	// intArr  := [...]int32{1,3,4}
	fmt.Printf("the intArray %v", intArr)

	var intSlice []int32 = []int32{2,3,4}
	fmt.Println(intSlice)

	var intSlice32 []int32 = []int32{8,9}
	var intSlice2 []int32 = []int32{5,6,3}

	intSlice = append(intSlice32, intSlice2...)
	fmt.Println(intSlice)

	//creates slice with initial length of 5 and capacity of 7
	var intSlice3 []int32 = make([]int32, 5, 7)

	//to add items in the slice
	intSlice3 = append(intSlice3, 10,20, 40,50)

	fmt.Printf("our intSlice3: %v \n", intSlice3)
	fmt.Printf("the len of slice is: %v \n",len(intSlice3))

	fmt.Println(intSlice3[1])

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint64 = map[string]uint64{"Adam":12, "Sarah": 15, "Serge":1000000000000, "Kanye":36787654563456}
	fmt.Println(myMap2)

	delete(myMap2, "Adam")

	var age, ok = myMap2["Adam"]

	if ok{
		fmt.Printf("the age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age:= range myMap2 {
		fmt.Printf("Name: %v, Value: %v \n", name, age)
	}

	//string in go use string library to process them
}