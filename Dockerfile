FROM golang:1.16.4 as builder

WORKDIR /go/src/github.com/afritzler/search-conversion
COPY . .

RUN make build-linux

FROM alpine:3.13.5

WORKDIR /

COPY --from=builder /go/src/github.com/afritzler/search-conversion/search-conversion_linux_amd64 /search-conversion

ENTRYPOINT ["/search-conversion"]