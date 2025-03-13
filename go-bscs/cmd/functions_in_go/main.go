package main

import (
	"errors"
	"fmt"
)

func main(){
	num1, num2 := 21, 5
	var result , remainder, err  = intDivision( num1, num2);

	returnDataIfConditionsCase(result, remainder, err);
	//returnDataRegularSwitchCase(result, remainder, err);
	//conditionalSwitchCase(result, remainder, err);
}


func returnDataRegularSwitchCase(result int, remainder int, err error){
	switch {
		case err != nil:
			fmt.Printf("%s", err.Error())
		case remainder == 0:
			fmt.Printf("the result is %v", result)
		default :fmt.Printf("the result of the division is %v and the remainder is %v", result, remainder)
	}
}


func conditionalSwitchCase(result int, remainder int, err error){

	switch remainder {
	case 0:
		fmt.Println("the division was exact")
	case 1,2:
		fmt.Println("the division was close")
	default:
		fmt.Println("the division was not close.")
	}
}

func returnDataIfConditionsCase(result int, remainder int, err error){
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if remainder == 0 {
		fmt.Printf("the result is %v", result)
	} else {
		fmt.Printf("the result of the division is %v and the remainder is %v", result, remainder)
	}
}

func intDivision(num1 int, num2 int) (int, int, error) {

	var err error

	if num2 == 0 {
		err = errors.New("cannot Divide by Zero")
		return 0, 0, err
	}

	return num1/num2, num1%num2, err
}

