package main

import (
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

	encodedCA, err := readCA(CAPath)
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

func readCA(caPath string) (string, error) {
	file, err := os.Open(CAPath)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	log.Printf("Data: %s", data)

	return string(data), nil
}
