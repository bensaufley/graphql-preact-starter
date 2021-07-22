FROM golang:1.16.6-alpine
LABEL maintainer="Ben Saufley<contact@bensaufley.com>"

SHELL ["/bin/ash", "-o", "pipefail", "-c"]
# hadolint ignore=DL3018
RUN apk add --no-cache --update curl sqlite=3.35.5-r0
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | \
  sh -s -- -b "$(go env GOPATH)/bin"

WORKDIR /go/src/github.com/bensaufley/graphql-preact-starter
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080
