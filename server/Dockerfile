FROM golang:1.17.1
LABEL maintainer="Ben Saufley<contact@bensaufley.com>"

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN apt-get update && \
  apt-get install -y --no-install-recommends sqlite3=3.34.1-3 && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | \
  sh -s -- -b "$(go env GOPATH)/bin"

WORKDIR /go/src/github.com/bensaufley/graphql-preact-starter
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080
