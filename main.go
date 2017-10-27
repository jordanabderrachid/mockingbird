package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/hashicorp/hcl"
)

type config struct {
	services []service
}

type service struct {
	name string
	host string
}

func main() {
	rawConfig, err := loadFile("./config.hcl")
	if err != nil {
		panic(err)
	}

	config := make(map[string]interface{})
	err = hcl.Decode(&config, rawConfig)
	if err != nil {
		panic(err)
	}

	services, err := extractConfig(config)
	if err != nil {
		panic(err)
	}

	router := createRouter(services)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func createRouter(services []service) *mux.Router {
	r := mux.NewRouter()

	for _, service := range services {
		route := r.NewRoute()
		route.Host(service.host)
		route.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(200)
		})
	}

	return r
}

func extractConfig(config map[string]interface{}) ([]service, error) {
	serviceValues, ok := config["service"]
	if !ok {
		return []service{}, nil
	}

	value := reflect.ValueOf(serviceValues)
	if value.Kind() != reflect.Slice {
		return []service{}, fmt.Errorf("invalid config, service must be an array. got type %s", value.Kind())
	}

	return extractServices(value.Index(0))
}

func extractServices(serviceConfigValue reflect.Value) ([]service, error) {
	if serviceConfigValue.Kind() != reflect.Map {
		return []service{}, fmt.Errorf("invalid service confing, must be a map. got type %s", serviceConfigValue.Kind())
	}

	services := make([]service, len(serviceConfigValue.MapKeys()))
	for i, serviceNameValue := range serviceConfigValue.MapKeys() {
		if serviceNameValue.Kind() != reflect.String {
			return []service{}, fmt.Errorf("invalid service confing, service name must be a string. got type %s", serviceNameValue.Kind())
		}

		services[i] = extractService(serviceNameValue.String(), serviceConfigValue.MapIndex(serviceNameValue))
	}

	return services, nil
}

func extractService(serviceName string, config reflect.Value) service {
	config = config.Elem()
	if config.Kind() != reflect.Slice {
		fmt.Printf("got type %s expected slice\n", config.Kind())
	}

	host := ""
	for _, configKey := range config.Index(0).MapKeys() {
		if configKey.String() == "host" {
			host = config.Index(0).MapIndex(configKey).Elem().String()
		}
	}

	return service{
		host: host,
		name: serviceName,
	}
}

func loadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return "", nil
	}

	return string(buf), nil
}
