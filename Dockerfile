FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/todo
COPY . .

RUN go install -v .

ENTRYPOINT ["/go/bin/todo"]
