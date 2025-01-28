package main

import "fmt"

func main() {
	multiple := 5 * 3
	quotient := multiple / 3

	isEven := multiple%2 == 0
	isGreater := multiple > quotient

	fmt.Println(multiple, quotient)
	fmt.Printf("is %d even? %t\n", multiple, isEven)
	fmt.Printf("is %d greater than %d? %t\n", multiple, quotient, isGreater)

	fmt.Printf("it is %t that %d is not greater than %d\n", !isGreater, multiple, quotient)
	// 														^^^^^^^^^^ not-logical operator (negation)
}
