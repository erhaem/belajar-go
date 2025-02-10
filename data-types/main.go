package main

import "fmt"

func main() {

	// i guess these variables' name and type
	// are pretty much self-explanatory
	// and giving a clue when to use each of them
	var username string = "erhaem"

	var smallNum int = 2
	var anotherSmallNum int32 = 200
	var bigNum int64 = 299999

	// use float32 for less precise measurements, like temperature
	var temperature float32 = 23.5 // celcius

	// use float64 for higher precision, like financial calculations
	var price float64 = 299.999

	var isGood bool = true

	fmt.Println("my username is", username)
	fmt.Println(smallNum, anotherSmallNum, bigNum)

	fmt.Printf("Temperature: %.2f°C\n", temperature)
	fmt.Printf("Price: $%.2f\n", price)

	fmt.Println("is even it good?", isGood)
}
