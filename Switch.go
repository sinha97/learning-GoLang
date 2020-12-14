package main

import (
	"fmt"
)
func main(){
	rating := 20;

	switch rating {

	case 15:
		fmt.Println("15")

	case 25:
		fmt.Println("25")
	
	default:
		fmt.Println("100")

	}

	temp := 10;

	switch  {

	case temp > 0:
		fmt.Println("temp is greater than zero")

	case temp == 0:
		fmt.Println("temp is zero")

	case temp < 0:
		fmt.Println("temp is smaller than zero")
		
	}
}