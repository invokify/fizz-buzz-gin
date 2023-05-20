package main

import (
	"fizz-buzz-gin/pkg/server"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// get the port from the environment variables
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}

	// get the port from the environment variables
	to, ok := os.LookupEnv("TIMEOUT")
	if !ok {
		to = "30"
	}

	// convert the timeout to integer
	timeout, err := strconv.Atoi(to)
	if err != nil {
		log.Printf("Failed to convert timeout to integer: %v", err)
		os.Exit(1)
	}

	log.Printf("Starting server on port %s", port)

	// instanciate the server
	s := server.NewServer(time.Duration(timeout) * time.Second)

	// start the server
	s.Run(":" + port)
}
