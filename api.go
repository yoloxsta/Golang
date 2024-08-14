package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting application...")

	// Intentional security issue: hardcoded API key
	apiKey := "AIzaSyD4u2f-SecretKeyThatShouldNotBeInCode"

	fmt.Println("Using API Key:", apiKey)
}
