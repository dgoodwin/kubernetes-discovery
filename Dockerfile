FROM golang

ADD . /go/src/github.com/dgoodwin/kubernetes-discovery
WORKDIR /go/src/github.com/dgoodwin/kubernetes-discovery
RUN go get && go install
ENTRYPOINT /go/bin/kubernetes-discovery

EXPOSE 8080
