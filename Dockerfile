FROM golang:alpine as BASE
WORKDIR /go/src/github.com/sul-dlss-labs/identifier-service
COPY . .
RUN apk update && \
    apk add --no-cache --virtual .build-deps git && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    apk del .build-deps
RUN CGO_ENABLED=0 GOOS=linux go build -o api -ldflags "-s" -a -installsuffix cgo cmd/app/main.go

FROM scratch
COPY --from=BASE /go/src/github.com/sul-dlss-labs/identifier-service/api .
CMD ["./api"]
