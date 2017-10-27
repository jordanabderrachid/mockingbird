package config

type Config struct {
	Services []Service
}

type Service struct {
	Name string
	Host string
}

// Parse reads the given string input and parses it as a Config struct.
func Parse(in string) (Config, error) {
	return Config{}, nil
}
