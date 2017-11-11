package config

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl"
)

// Server represents the configuration of the server.
type Server struct {
	Services []Service `hcl:"service"`
}

// Service represents one mocked service.
type Service struct {
	Name      string     `hcl:",key"`
	Host      string     `hcl:"host"`
	Endpoints []Endpoint `hcl:"endpoint,expand"`
}

// Endpoint .
type Endpoint struct {
	Name   string `hcl:",key"`
	Method string `hcl:"method"`
	Path   string `hcl:"path"`
}

func (s Server) String() string {
	services := make([]string, len(s.Services))
	for i, service := range s.Services {
		services[i] = fmt.Sprintf("\t%s", service.String())
	}

	return fmt.Sprintf("[services]\n%s", strings.Join(services, "\n"))
}

func (s Service) String() string {
	endpoints := make([]string, len(s.Endpoints))
	for i, endpoint := range s.Endpoints {
		endpoints[i] = fmt.Sprintf("\t\t%s", endpoint.String())
	}

	return fmt.Sprintf("[service] name=%s, host=%s\n%s", s.Name, s.Host, strings.Join(endpoints, "\n"))
}

func (e Endpoint) String() string {
	return fmt.Sprintf("[endpoint] name=%s, method=%s, path=%s", e.Name, e.Method, e.Path)
}

// Parse reads the given string input and parses it as a Config struct.
func Parse(in string) (*Server, error) {
	server := &Server{}
	err := hcl.Decode(server, in)
	return server, err
}
