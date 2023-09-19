// Package denv manages whether the application is running in Development, Integration, Staging, or Production.
package denv

import (
	"fmt"
	"os"
)

const (
	// Development is the development environment.
	Development = "development"
	// Integration is the integration environment.
	Integration = "integration"
	// Staging is the staging environment.
	Staging = "staging"
	// Production is the production environment.
	Production = "production"
)

// Env is the environment the application is running in.
// It reads the configuration from the environment variable DENV_DEPLOYMENT_ENV.
// This can be set to "development", "integration", "staging", or "production".
type Env struct {
	// Environment the application is running in.
	// This can be set to "development", "integration", "staging", or "production".
	Environment string
}

// NewEnv creates a new Env.
// If the environment variable DENV_DEPLOYMENT_ENV is not set, it returns an ErrInvalidEnvironment.
func NewEnv() (*Env, error) {
	e := Env{}
	e.Environment = os.Getenv("DENV_DEPLOYMENT_ENV")

	if e.Environment != Development && e.Environment != Integration && e.Environment != Staging && e.Environment != Production {
		return nil, fmt.Errorf("%w: %s", ErrInvalidEnvironment, e.Environment)
	}
	return &e, nil
}

// String returns the environment as a string.
func (e *Env) String() string {
	return e.Environment
}

// IsDevelopment returns true if the environment is "development".
func (e *Env) IsDevelopment() bool {
	return e.Environment == "development"
}

// IsIntegration returns true if the environment is "integration".
func (e *Env) IsIntegration() bool {
	return e.Environment == "integration"
}

// IsStaging returns true if the environment is "staging".
func (e *Env) IsStaging() bool {
	return e.Environment == "staging"
}

// IsProduction returns true if the environment is "production".
func (e *Env) IsProduction() bool {
	return e.Environment == "production"
}
