FROM golang:1.16-alpine as go-builder

WORKDIR /go/src/github.com/bensaufley/graphql-preact-starter
COPY go.mod go.sum ./
RUN go mod download

COPY server/* ./

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix 'static' -o server ./cmd/server

FROM node:16.5.0 as node-builder

WORKDIR /tmp
COPY package*.json ./
RUN npm install

WORKDIR /usr/src/graphql-preact-starter
RUN mv /tmp/package*.json /tmp/node_modules ./

COPY client/* ./

ENV NODE_ENV=production

RUN npm run build

RUN npm prune --production

FROM alpine:3.14

WORKDIR /app
COPY --from=go-builder /go/src/github.com/bensaufley/graphql-preact-starter/server .
COPY --from=node-builder /usr/src/graphql-preact-starter/.build /public
