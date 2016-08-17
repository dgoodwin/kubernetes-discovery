package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TODO:
const tempToken string = "mytoken"

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

	clusterInfo := ClusterInfo{
		Type:             "ClusterInfo",
		Version:          "v1",
		RootCertificates: "NOTHINGHEreYET!",
	}

	if err := json.NewEncoder(resp).Encode(clusterInfo); err != nil {
		panic(err)
	}
}
