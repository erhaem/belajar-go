package main

import "fmt"

func main() {

	// this following called "type manifesting" variables -- the explicit way
	// syntax: var <variable-name> <data-type>
	var name string = "erhaem"
	// you can declare first, assign later :p
	var age int
	age = 420

	// in case you're unsure, go with "type inference" -- the implicit way
	// let the compiler decide the type based on the assigned value
	// you don't need to write "var <variable-name> <data-type>"
	// instead, "<variable-name> := <value>"
	description := "i'm currently learning go. and i write this code using nano"

	fmt.Printf("my name is %s, a %d years old. %s\n", name, age, description)

	// multiple varirable declarations
	var one, two, three string = "one", "two", "three"
	four, five, six := 4, 5, "six"

	fmt.Println(one, two, three, four, five, six)

	// underscore "variable" -- actually not a variable
	// in case a variable potentially won't be used,
	// mostly used for ignoring function return values, like errors
	_, c := "john", "wick" // let's just say "john" is an error
	fmt.Println(c)
	//fmt.Println(_)
	// printing _ causes error, because it's a blank identifier, a placeholder
	// it yells at the compiler
	// "don't assign me a value - i won't store and use it anyway!"

	// using new()
	// i still don't get it -- has to do with pointer
	// and i don't understand pointer right now
	// TODO update this later
	isSleep := new(bool)
	fmt.Println(isSleep)
	fmt.Println(*isSleep)
}
