package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// CAPath is the expected location of our cluster's CA to be distributed to
// clients looking to connect. Because we expect to use kubernetes secrets
// for the time being, this file is expected to be a base64 encoded version
// of the normal cert PEM.
const CAPath = "/tmp/secret/ca.pem.b64"

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

	// TODO: Just verifying our cert is decoding properly, this can all
	// be removed.
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("ERROR: Unable to read %s", CAPath)
	}
	file.Close()
	b64Str := string(data)
	decodedPEM, err := base64.StdEncoding.DecodeString(b64Str)
	log.Printf("Decoded PEM: \n\n%s", decodedPEM)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
