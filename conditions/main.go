package main

import "fmt"

func main() {
	score := 90

	fmt.Println("your score is", score)

	// going with if statements (w/ temp variable)
	if diff := 100 - score; score > 100 {
		fmt.Println("ooof, maximum score is 100")
	} else if score == 100 {
		fmt.Println("perfect score lel")
	} else if score >= 50 {
		fmt.Printf("not bad :/ (you missed %d points)\n", diff)
	} else if score >= 10 {
		fmt.Printf("a monke can do this better, cmon (you missed %d points)\n", diff)
	} else if score == 0 {
		fmt.Println("WTAF were ya doing all these times?!")
	} else {
		fmt.Println("who r ya?")
	}

	// switch-case statement w/ if-statement style
	// because the regular one is boring
	var scoreLabel string
	switch {
	case score == 100:
		scoreLabel = "A"
	case score > 50:
		scoreLabel = "B"
	case score > 10:
		scoreLabel = "C"
	default:
		scoreLabel = "D"
	}

	fmt.Println("it is", scoreLabel)
}
