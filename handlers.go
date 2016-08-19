package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO: Just using a hardcoded token for now.
const tempToken string = "TOKENID.TOKEN"

func Index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Welcome!")
}

func ClusterInfoIndex(resp http.ResponseWriter, req *http.Request) {

	tokenId := req.FormValue("token-id")
	log.Printf("Got token ID: %s", tokenId)
	if tokenId != tempToken {
		log.Print("Invalid token")
		http.Error(resp, "Forbidden", http.StatusForbidden)
		return
	}

	encodedCA, err := readAndEncodeCA(CAPath)
	if err != nil {
		http.Error(resp, "Error encoded CA", http.StatusInternalServerError)
		return
	}

	clusterInfo := ClusterInfo{
		Type:             "ClusterInfo",
		Version:          "v1",
		RootCertificates: encodedCA,
	}

	if err := json.NewEncoder(resp).Encode(clusterInfo); err != nil {
		panic(err)
	}
}

func readAndEncodeCA(caPath string) (string, error) {
	file, err := os.Open(CAPath)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	log.Printf("Data: %s", data)

	encodedCA := base64.StdEncoding.EncodeToString([]byte(data))
	log.Printf("Encoded: %s", encodedCA)
	return encodedCA, nil
}
