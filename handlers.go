package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/square/go-jose.v1"
)

// TODO: Just using a hardcoded token for now.
const tempTokenId string = "TOKENID"
const tempToken string = "EF1BA4F26DDA9FE2"

func Index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Welcome!")
}

// ClusterInfoHandler implements the http.ServeHTTP method and allows us to
// mock out portions of the request handler in tests.
type ClusterInfoHandler struct {
}

func (cih *ClusterInfoHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	tokenId := req.FormValue("token-id")
	log.Printf("Got token ID: %s", tokenId)
	if tokenId != tempTokenId {
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

	// Instantiate an signer using HMAC-SHA256.
	hmacTestKey := fromHexBytes(tempToken)
	signer, err := jose.NewSigner(jose.HS256, hmacTestKey)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Error creating JWS signer: %s", err), http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(clusterInfo)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Error serializing clusterInfo to JSON: %s", err),
			http.StatusInternalServerError)
		return
	}

	// Sign a sample payload. Calling the signer returns a protected JWS object,
	// which can then be serialized for output afterwards. An error would
	// indicate a problem in an underlying cryptographic primitive.
	jws, err := signer.Sign(payload)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Error signing clusterInfo to JSON: %s", err),
			http.StatusInternalServerError)
		return
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized := jws.FullSerialize()

	resp.Write([]byte(serialized))

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

	encodedCA := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedCA, nil
}

// TODO: Move into test package
// TODO: Should we use base64 instead?
func fromHexBytes(base16 string) []byte {
	val, err := hex.DecodeString(base16)
	if err != nil {
		panic(fmt.Sprintf("Invalid test data: %s", err))
	}
	return val
}
