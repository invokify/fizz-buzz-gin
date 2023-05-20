package main

import (
	"fizz-buzz-gin/pkg/server"
	"log"
	"os"
)

func main() {
	// get the port from the environment variables
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}

	log.Printf("Starting server on port %s", port)

	// instanciate the server
	s := server.NewServer()

	// start the server
	s.Run(":" + port)
}
