#FROM golang:alpine AS builder
#WORKDIR /build
#COPY . .
#RUN go build -o build main.go

FROM alpine
WORKDIR /build
COPY /build/main /build
CMD ["/build/main"]