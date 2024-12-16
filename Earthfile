# Earthfile
VERSION 0.8
# Use the base image built in Earthfile.base
# FROM +base:build AS base
FROM golang:1.23-alpine
WORKDIR /go-example


deps:
    # FROM base
    COPY go.mod ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    # SAVE ARTIFACT go.sum AS LOCAL go.sum

test:
    FROM +deps
    COPY . .
    RUN go test ./...

verify:
    ARG registry
    RUN echo "registry: $registry"
    ARG buildversion
    ARG binaryDefault
    ARG dockerbin
    RUN echo "buildversion: $buildversion"
    RUN echo "------"


build:
    FROM +deps
    COPY . .
    ARG buildversion
    RUN echo "buildversion: $buildversion"
    RUN go build -o bin/go-example-$buildversion cmd/site/main.go
    SAVE ARTIFACT bin/go-example-$buildversion AS LOCAL bin/go-example-$buildversion

docker:
    FROM +build
    # WORKDIR /go-example
    ARG buildversion
    ARG registry
    RUN apk add --no-cache curl

    COPY +build/go-example-$buildversion ./go-example
    COPY static ./static
    ENTRYPOINT ["./go-example"]
    SAVE IMAGE --push $registry/earthly:$buildversion
    SAVE IMAGE --push $registry/earthly:latest


docker-test:
    FROM earthly/examples:go
    WORKDIR /
    RUN apk add --no-cache curl
    # Start the app in the background
    RUN ["sh", "-c", "./go-example & sleep 2"]
    # Test if the website responds with HTTP 200
    RUN curl -I http://localhost:8080 | grep "HTTP/1.1 200 OK"
    
all:
    BUILD +test
    BUILD +build
    BUILD +docker
