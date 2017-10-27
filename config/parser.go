package config

import (
	"strings"
	"fmt"

	"github.com/hashicorp/hcl"
)

// Config represents the configuration of the server.
type Config struct {
	Services []Service
}

func (c Config) String() string {
	services := make([]string, len(c.Services))
	for i, service := range c.Services {
		services[i] = fmt.Sprintf("\t%s", service.String())
	}

	return fmt.Sprintf("[services]\n%s", strings.Join(services, "\n"))
}

// Service represents one mocked service.
type Service struct {
	Name string
	Host string
}

func (s Service) String() string {
	return fmt.Sprintf("[service] name=%s, host=%s", s.Name, s.Host)
}

// Parse reads the given string input and parses it as a Config struct.
func Parse(in string) (Config, error) {
	out := make(map[string]interface{})
	err := hcl.Decode(out, in)
	if err != nil {
		return Config{}, err
	}

	return Config{}, nil
}
