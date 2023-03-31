package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Serve files from static folder
	http.Handle("/", http.FileServer(http.Dir("/app")))

	port := ":8000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}
