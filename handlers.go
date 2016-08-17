package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ClusterInfoIndex(w http.ResponseWriter, r *http.Request) {
	clusterInfo := ClusterInfo{
		Type:             "ClusterInfo",
		Version:          "v1",
		RootCertificates: "NOTHINGHEreYET!",
	}

	if err := json.NewEncoder(w).Encode(clusterInfo); err != nil {
		panic(err)
	}
}
