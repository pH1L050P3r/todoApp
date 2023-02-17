FROM golang:1.18-alpine

WORKDIR /project
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY ./pkg/ ./pkg/
COPY ./cmd/ ./cmd/
COPY ./.env ./.env

RUN go build -o ./build/app.o ./cmd/web/*.go

EXPOSE 8000