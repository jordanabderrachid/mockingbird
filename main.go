package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jordanabderrachid/mockingbird/config"
)

func main() {
	rawConfig, err := loadFile("./config.hcl")
	if err != nil {
		panic(err)
	}

	server, err := config.Parse(rawConfig)
	if err != nil {
		panic(err)
	}

	log.Printf("server listening on :8080\n%s\n", server)
	router := createRouter(server)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func createRouter(server *config.Server) *mux.Router {
	r := mux.NewRouter()

	for _, service := range server.Services {
		for _, endpoint := range service.Endpoints {
			registerEndpoint(r, service.Host, endpoint)
		}
	}

	return r
}

func registerEndpoint(router *mux.Router, hostname string, endpoint config.Endpoint) {
	for _, behavior := range endpoint.Behaviors {
		router.HandleFunc(endpoint.Path, translateBehavior(behavior)).Host(hostname).Methods(endpoint.Method)
	}
}

func translateBehavior(behavior config.Behavior) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(behavior.Response.Code)
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
