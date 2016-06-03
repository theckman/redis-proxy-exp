package main

import (
	"fmt"
	"net/http"
	"sync"

	"gopkg.in/redis.v3"
)

// Backend is the representation of the Redis backend. It includes the Client,
// and it implements the http.Handler interface.
type Backend struct {
	Addr string

	client   *redis.Client
	clientMu sync.Mutex
}

func (b *Backend) prepBackend() error {
	b.clientMu.Lock()
	defer b.clientMu.Unlock()

	if b.client == nil {
		b.client = redis.NewClient(&redis.Options{
			Addr: b.Addr,
		})

		// do a ping to ensure we can contact the server
		_, err := b.client.Ping().Result()

		return err
	}

	return nil
}

func (b *Backend) get(key string) ([]byte, error) {
	// generate a new Redis client if we don't have one
	if b.client == nil {
		err := b.prepBackend()

		if err != nil {
			return nil, fmt.Errorf("error prepping backend: %s\n", err)
		}
	}

	reply := b.client.Get(key)

	if err := reply.Err(); err != nil {
		return nil, err
	}

	return reply.Bytes()
}

func getKey(r *http.Request) string {
	param := r.URL.Query()["key"]

	if len(param) < 1 {
		return ""
	}

	return param[0]
}

func (b *Backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// set the Content-Type to text/plain
	w.Header().Add("content-type", "text/plain")

	// get the key from the query params
	// will be "" if the expected key is not present
	key := getKey(r)

	// if the key is not provided, it's a 400 error
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "you must provide a 'key' GET parameter\n")
		return
	}

	// let's try to get the data from Redis
	value, err := b.get(key)

	// if there's an error, we need to return it to the client
	if err != nil {
		// redis.Nil is returned if the key is not found
		// so throw a 404
		if err == redis.Nil {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "key '%s' not found", key)
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintf(w, "error retrieving key '%s': %s\n", key, err)
		}

		return
	}

	w.Write(value)
}
