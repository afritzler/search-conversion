FROM golang:1.18.5 as builder

WORKDIR /go/src/github.com/afritzler/search-conversion
COPY . .

RUN make build-linux

FROM alpine:3.16.1

WORKDIR /

COPY --from=builder /go/src/github.com/afritzler/search-conversion/search-conversion_linux_amd64 /search-conversion

ENTRYPOINT ["/search-conversion"]