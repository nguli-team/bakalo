package config

import "fmt"

// Environment is application deployment type
type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
)

func (e Environment) IsValid() error {
	switch e {
	case Production, Development:
		return nil
	}
	return fmt.Errorf("environment configuration is invalid: \"%s\"\n", e)
}
