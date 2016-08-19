package main

import (
	"log"
	"net/http"
	"os"
)

// CAPath is the expected location of our cluster's CA to be distributed to
// clients looking to connect. Because we expect to use kubernetes secrets
// for the time being, this file is expected to be a base64 encoded version
// of the normal cert PEM.
const CAPath = "/tmp/secret/ca.pem"

func main() {

	// Make sure the CA cert for the cluster exists and is readable.
	// We are expecting a base64 encoded version of the cert PEM as this is how
	// the cert would most likely be provided via kubernetes secrets.
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
	log.Printf("Listening for requests on port 9898.")
	log.Fatal(http.ListenAndServe(":9898", router))
}
