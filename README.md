# kubernetes-discovery
An initial implementation of a Kubernetes discovery service using JSON Web Signatures.

This prototype is expected to be run by Kubernetes itself for the time being,
and will hopefully be merged into the core API at a later time.

## Running from source
- Generate a CA cert and base64 encode it, saving the result to /tmp/secret/ca.pem.b64.
- `go get`
- `go install && ~/go/bin/kubernetes-discovery`

## Running in Docker
- `docker build -t kubernetes-discovery-proto .`
- `docker run -p 8080:8080 -v /home/dgoodwin/go/src/github.com/dgoodwin/kubernetes-discovery/ca.pem:/var/lib/kubernetes/ca.pem --name kubernetes-discovery --rm kubernetes-discovery-proto`

## Running in Kubernetes

TODO

## Testing the API

`curl "http://localhost:8080/cluster-info/v1/?token-id=TOKENID.TOKEN"`



