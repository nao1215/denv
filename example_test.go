package denv_test

import (
	"fmt"
	"os"

	"github.com/nao1215/denv"
)

func Example() {
	if err := os.Setenv("APP_ENV", denv.Staging); err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	// Create a new Env instance.
	env, err := denv.NewEnv()
	if err != nil {
		fmt.Println("Error creating environment:", err)
		return
	}

	// Print the current environment.
	fmt.Printf("Current environment: %s\n", env.String())

	// Check if it's the development environment.
	if env.IsDevelopment() {
		fmt.Println("This is the development environment.")
	} else {
		fmt.Println("This is not the development environment.")
	}

	// Check if it's the production environment.
	if env.IsProduction() {
		fmt.Println("This is the production environment.")
	} else {
		fmt.Println("This is not the production environment.")
	}

	// Output:
	// Current environment: staging
	// This is not the development environment.
	// This is not the production environment.
}
