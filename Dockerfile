FROM golang:1.18-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /project

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/app.o ./cmd/web/*.go

FROM alpine:latest

WORKDIR /project/

COPY --from=builder /project/build/app.o .
COPY --from=builder /project/.env .

EXPOSE 8000

ENTRYPOINT ["/bin/sh", "-c"]