FROM golang:1.23rc1-alpine3.20

RUN apk update && apk add git

WORKDIR /go/src/github.com/organization-ranking

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . ./

RUN go install golang.org/x/tools/cmd/goimports@latest
