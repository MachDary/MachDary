# Build in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make git

ADD . /go/src/github.com/MachDary/MachDary
RUN cd /go/src/github.com/MachDary/MachDary && make server && make client

# Pull into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/MachDary/MachDary/cmd/server/server /usr/local/bin/
COPY --from=builder /go/src/github.com/MachDary/MachDary/cmd/client/client /usr/local/bin/

EXPOSE 1999 60516 60517 6051
