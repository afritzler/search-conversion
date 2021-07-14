FROM golang:1.17rc1 as builder

WORKDIR /go/src/github.com/afritzler/search-conversion
COPY . .

RUN make build-linux

FROM alpine:3.14.0

WORKDIR /

COPY --from=builder /go/src/github.com/afritzler/search-conversion/search-conversion_linux_amd64 /search-conversion

ENTRYPOINT ["/search-conversion"]