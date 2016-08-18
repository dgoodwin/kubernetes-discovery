package main

import (
	"log"
	"net/http"
	"os"
)

// CAPath is the expected location of our cluster's CA to be distributed to
// clients looking to connect.
const CAPath = "/var/lib/kubernetes/ca.pem"

func main() {

	// Make sure the CA cert for the cluster exists and is readable:
	if _, err := os.Stat(CAPath); os.IsNotExist(err) {
		log.Fatalf("CA does not exist: %s", CAPath)
	}
	// Test read permissions
	file, err := os.Open(CAPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to read %s", CAPath)
	}
	file.Close()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
