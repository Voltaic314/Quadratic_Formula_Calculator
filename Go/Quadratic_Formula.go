package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var a_str string
	var b_str string
	var c_str string

	fmt.Println("Hello World!")
	fmt.Println("Today we will be calculating the roots of X given your A, B, and C values.")
	
	fmt.Printf("First, give us your A value: ")
	fmt.Scanln(&a_str)

	a, err := strconv.ParseFloat(a_str, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Next, give us your B value: ")
	fmt.Scanln(&b_str)

	b, err := strconv.ParseFloat(b_str, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Finally, give us your C value: ")
	fmt.Scanln(&c_str)
	
	c, err := strconv.ParseFloat(c_str, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Great! Now we will calculate the roots of x. Give us one moment...")

	discriminant := (b * b) - (4 * a * c) // Go will infer the type (float64) automatically
	if discriminant < 0 {
		fmt.Println("Your formula has no real solutions.")
		return
	}

	x_1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	x_2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	fmt.Printf("Your solutions are: %f and %f", x_1, x_2)
	fmt.Scanln()
}
