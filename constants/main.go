package main

import "fmt"

func main() {
	//currently having my eyes on https://blog.boot.dev/clean-code/constants-in-go-vs-javascript-and-when-to-use-them/

	// below is a constant, once you declared it u can't modify the value
	const firstName, lastName string = "Agus", "Susanto"

	const schoolName string = "SMKN 666"

	// this wont work:
	// firstName = "Agung"
	// once "Agus" always "Agus"!

	// to declare constants, you can also do this:
	const (
		age  int = 29
		city     = "New York" // this one use type inference
	)

	// here is ugly output
	fmt.Println(
		"My name is", firstName+" "+lastName,
		".", age, "is my age", // does this remind you of something?
		". I live in ", city,
		". I go to a boring school daily, "+schoolName,
	)
}
