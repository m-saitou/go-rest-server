FROM alpine:latest

WORKDIR /usr/local/bin/

COPY gopath/bin/go-rest-server .

ENTRYPOINT ["/usr/local/bin/go-rest-server"]