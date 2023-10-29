// This goddamn code is temporary!! If you're still reading this, get off your ass and add more shit!!

package main

import (
	"fmt"
)

func main() {
	//prints init message
	fmt.Println("insert msg (treat underscores as spaces lol) ")
	// number and read input
	var inputtest string
	fmt.Scan(&inputtest)
	//math operation to complete
	var operator string
	fmt.Scan(&operator)
	//operate by this number
	var inputtest3 string
	fmt.Scan(&inputtest3)
	if operator == "+" {
		operator = inputtest + inputtest3
	}
	fmt.Println(operator)

}
