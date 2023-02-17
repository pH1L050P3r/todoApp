FROM golang:1.16-alpine

WORKDIR /project
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY ./pkg/ ./pkg/
COPY ./cmd/ ./cmd/

RUN go build -o ./build/app ./cmd/web/*.go

EXPOSE 8000