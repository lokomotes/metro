FROM golang:1.11 AS golang

COPY . /go/.

RUN go build -ldflags "-linkmode external -extldflags -static" -a \
    /go/src/github.com/lokomotes/metro/cmd/metro-server/main.go



FROM scratch

COPY --from=golang /go/main /metro-server

ENTRYPOINT [ "/metro-server" ]
