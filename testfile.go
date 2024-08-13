package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting application...")

	// Intentional compilation error: undeclared variable
	result := addNumbers(5, 10)
	fmt.Println("Result:", result)
}

// Function is missing
// func addNumbers(a int, b int) int {
//     return a + b
// }
