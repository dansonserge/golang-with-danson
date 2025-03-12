package main

import "fmt"

func main(){
	var intNum int = 8977
	var intNum16 int16 = 8977
	var intNum32 int32 = 345345
	var intNum64 int64 = 3453452345234523
	fmt.Println(intNum, intNum16, intNum32, intNum64)

	var floatNum32 , floatNum64= 234.4, 345.7

	fmt.Println(floatNum32, floatNum64)

	var names string = "wertwer"
	phrase := `I am
		testing
	`
	fmt.Println(phrase)

	fmt.Println(len(names))

	var myRune rune = 'A'
	fmt.Println(myRune)

	town1, town2, town3:= "Brooklyn", "Queens", "Staten Island"

	fmt.Println(town1, town2, town3)

	const pi float32 = 3.1415
	fmt.Println(pi)

}