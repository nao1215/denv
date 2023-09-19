![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/denv/coverage.svg)
[![LinuxUnitTest](https://github.com/nao1215/denv/actions/workflows/linux_test.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/nao1215/denv/actions/workflows/mac_test.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/nao1215/denv/actions/workflows/windows_test.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/windows_test.yml)
[![FreeBSDTest](https://github.com/nao1215/denv/actions/workflows/freebsd.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/freebsd.yml)
[![NetBSDTest](https://github.com/nao1215/denv/actions/workflows/netbsd.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/netbsd.yml)
[![OpenBSDTest](https://github.com/nao1215/denv/actions/workflows/openbsd.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/openbsd.yml)
[![DragonflyBSDTest](https://github.com/nao1215/denv/actions/workflows/dragonfly.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/dragonfly.yml)
[![Gosec](https://github.com/nao1215/denv/actions/workflows/security.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/security.yml)
[![Vuluncheck](https://github.com/nao1215/denv/actions/workflows/govulncheck.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/govulncheck.yml)
[![reviewdog](https://github.com/nao1215/denv/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/denv/actions/workflows/reviewdog.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/denv.svg)](https://pkg.go.dev/github.com/nao1215/denv)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/denv)](https://goreportcard.com/report/github.com/nao1215/denv)
# denv
"denv" is a library that manages the deployment environment, which can be one of development, integration, staging, or production.

## How to use
1. You import the library: "github.com/nao1215/denv"
2. You set the environment variable "DENV_DEPLOYMENT_ENV" to one of "development", "integration", "staging", or "production".
3. You create a new Env instance with denv.NewEnv().
   
```go
func Example() {
	if err := os.Setenv("DENV_DEPLOYMENT_ENV", denv.Staging); err != nil {
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