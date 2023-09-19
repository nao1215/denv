# denv
"denv" is a library that manages the deployment environment, which can be one of development, integration, staging, or production.

## How to use
1. You import the library: "github.com/nao1215/denv"
2. You set the environment variable "DENV_DEPLOYMENT_ENV" to one of "development", "integration", "staging", or "production".
3. You create a new Env instance with denv.NewEnv().
   
```go
func Example() {
	os.Setenv("DENV_DEPLOYMENT_ENV", denv.Staging)

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
```

## Support OS & go version
- Linux
- Mac
- Windows
- FreeBSD
- OpenBSD
- NetBSD
- DragonflyBSD
- go version 1.20 or later

## LICENSE
[MIT](./LICENSE)
```