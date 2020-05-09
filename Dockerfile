FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/todo
COPY . .

RUN go install -v .

ENTRYPOINT ["sleep", "10000"]
# ENTRYPOINT ["/go/bin/todo"]
