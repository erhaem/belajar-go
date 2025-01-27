package main

import "fmt"

func main() {
	multiple := 5 * 3
	iGotNoIdea := multiple / 3

	isEven := multiple%2 == 0
	isGreater := multiple > iGotNoIdea

	fmt.Println(multiple, iGotNoIdea)
	fmt.Println("is", multiple, "even?", isEven)
	fmt.Println("is", multiple, "greater than", iGotNoIdea, "?", isGreater)

	// this quite boring. i quit
}
