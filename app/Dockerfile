FROM golang:1.21-bullseye AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod ./
COPY vendor ./vendor
COPY . ./
RUN go build -mod=vendor -o server ./main.go

FROM debian:9

WORKDIR /app
COPY --from=builder /app/server /app/server

EXPOSE 8080
CMD ["/app/server"]
