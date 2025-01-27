package main

import (
	"os"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Give us your A, B, and C values as comma-separated values: ")

	// Read the full input line
	user_input, _ := reader.ReadString('\n')
	user_input = strings.TrimSpace(user_input) // Remove extra spaces or newlines
	fmt.Printf("Raw input: %q\n", user_input)

	// Split the input into parts
	split_str := strings.Split(user_input, ",")
	fmt.Printf("Split result: %#v\n", split_str)

	if len(split_str) != 3 {
		fmt.Println("Please ensure you are entering 3 values.")
		return
	}

	a, a_err := strconv.ParseFloat(strings.TrimSpace(split_str[0]), 64)
	b, b_err := strconv.ParseFloat(strings.TrimSpace(split_str[1]), 64)
	c, c_err := strconv.ParseFloat(strings.TrimSpace(split_str[2]), 64)
	if a_err != nil || b_err != nil || c_err != nil {
		fmt.Println("Please ensure you are entering valid numbers.")
		return
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
