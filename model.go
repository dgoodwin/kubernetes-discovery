package main

type ClusterInfo struct {
	Type             string
	Version          string
	RootCertificates string `json:"rootCertificates"`
	// TODO: ClusterID, Endpoints
}
