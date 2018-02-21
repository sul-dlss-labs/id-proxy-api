FROM golang:alpine
WORKDIR /go/src/github.com/sul-dlss-labs/identifier-service
COPY . .
RUN apk update && \
    apk add --no-cache --virtual .build-deps git && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    apk del .build-deps

WORKDIR /go/src/github.com/sul-dlss-labs/identifier-service/cmd/app
RUN go install .

CMD ["app"]
