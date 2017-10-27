package config_test

import "testing"
import "github.com/jordanabderrachid/mockingbird/config"
import "github.com/stretchr/testify/assert"

func TestParse(t *testing.T) {
	var cases = []struct {
		in     string
		config config.Config
	}{
		{
			in: `
			service "greeter" {
				host = "greeter.service"
			}
			`,
			config: config.Config{
				Services: []config.Service{config.Service{Host: "greeter.service", Name: "greeter"}},
			},
		},
	}

	for _, c := range cases {
		actualConfig, err := config.Parse(c.in)
		if err != nil {
			t.Fatalf("Fail to parse config. in=%s, expected=%s, err=%s",
				c.in, c.config, err.Error())
		}

		assert.Equal(t, c.config, actualConfig)
	}
}
