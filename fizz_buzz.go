package algorithms

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		fizzBuzzValue(i)
		fmt.Print(", ")
	}
	fizzBuzzValue(n)
	fmt.Println()
}

func fizzBuzzValue(i int) {
	switch {
	case i%3 == 0 && i%5 == 0:
		fmt.Printf("Fizz Buzz")
	case i%5 == 0:
		fmt.Printf("Buzz")
	case i%3 == 0:
		fmt.Printf("Fizz")
	default:
		fmt.Print(i)
	}
}
