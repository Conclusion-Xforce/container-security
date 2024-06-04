/*
 Simple golang REST service to connect to a redis service and store a key value pair
 The service will also have a health check endpoint to check the status of the service

 The service exposes the following endpoints:
   /health 		 	  | GET  		| returns the status of the service			
   /store/KEY 		  | GET  		| retrieves the value of the key from the redis service			
   /store/KEY/VALUE   | POST 		| stores a key value pair in the redis service			

 The service needs the following environment variables to be set:
   SERVICE_PORT - the port on which the service will run (defaults to 80)
   REDIS_HOST 	- the hostname of the redis service
   REDIS_PORT 	- the port of the redis service

 The service logs all requests to valid endpoints
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/redis/go-redis/v9"
	"github.com/gorilla/mux"
)

var client *redis.Client
var ctx = context.Background()

func initRedisClient() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func main() {
	initRedisClient()

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/store/{key}", retrieveHandler).Methods("GET")
	r.HandleFunc("/store/{key}/", retrieveHandler).Methods("GET")
	r.HandleFunc("/store/{key}/{value}", storeHandler).Methods("POST")

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
}

func logRequest(r *http.Request, statusCode int) {
	log.Printf("%s %s %s %s %d", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, statusCode)
}

func writeResponse(w http.ResponseWriter, statusCode int, response string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response + "\n"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := client.Ping(ctx).Result()

	if err != nil {
		logRequest(r, http.StatusInternalServerError)
		writeResponse(w, http.StatusInternalServerError, "Failed to connect to Redis")
		return
	}
	logRequest(r, http.StatusOK)
	writeResponse(w, http.StatusOK, "Service is up")
}

func storeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	// Store the key value pair in the redis service
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		logRequest(r, http.StatusInternalServerError)
		writeResponse(w, http.StatusInternalServerError, "Failed to store " + key)
		return
	}

	logRequest(r, http.StatusOK)
	writeResponse(w, http.StatusOK, value)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// Retrieve the value of the key from the redis service
	value, err := client.Get(ctx, key).Result()

	// If the key does not exist, return a 404
	if err == redis.Nil {
		logRequest(r, http.StatusNotFound)
		writeResponse(w, http.StatusNotFound, "Key not found")
		return
	}

	if err != nil {
		logRequest(r, http.StatusInternalServerError)
		writeResponse(w, http.StatusInternalServerError, "Failed to retrieve " + key)
		return
	}

	logRequest(r, http.StatusOK)
	writeResponse(w, http.StatusOK, value)
}
