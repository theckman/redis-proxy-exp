package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const fciVersion = "0.0.2"

func buildMux(b *Backend) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/get", b)
	return router
}

func getArgs() (*binArgs, error) {
	args := &binArgs{}

	out, err := args.parse(nil)

	if err != nil {
		return nil, err
	}

	// if out is not empty we should print it and exit 0
	if out != "" {
		fmt.Print(out)
		os.Exit(0)
	}

	return args, nil
}

func main() {
	args, err := getArgs()

	if err != nil {
		log.Fatalf("error parsing arguments: %s", err)
	}

	b := &Backend{Addr: fmt.Sprintf("%s:%d", args.RedisHost, args.RedisPort)}

	// build the router and serve the requests
	router := buildMux(b)
	http.ListenAndServe(fmt.Sprintf(":%d", args.Port), router)
}
