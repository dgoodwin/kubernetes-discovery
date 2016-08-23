# kubernetes-discovery
An initial implementation of a Kubernetes discovery service using JSON Web Signatures.

This prototype is expected to be run by Kubernetes itself for the time being,
and will hopefully be merged into the core API at a later time.

## Running from source
- Generate a CA cert save it to: /tmp/secret/ca.pem
- `go get`
- `go install && ~/go/bin/kubernetes-discovery`

## Running in Docker
- `docker run --rm -p 9898:9898 -v /tmp/secret/ca.pem:/tmp/secret/ca.pem --name kube-disco dgoodwin/kube-disco`

## Running in Kubernetes

A dummy certificate is included in ca-secret.yaml.

```
create -f ca-secret.yaml
create -f kube-disco.yaml
```

## Testing the API

`curl "http://localhost:9898/cluster-info/v1/?token-id=TOKENID"`



