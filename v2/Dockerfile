FROM golang:1.22-alpine as builder

WORKDIR $GOPATH/src/github.com/saucon/sauron
COPY . .
RUN ls -al


WORKDIR $GOPATH/src/github.com/saucon/sauron/v2

RUN go build -o sauron .
RUN ls -al

FROM alpine:latest

COPY --from=builder /go/src/github.com/saucon/sauron/v2/sauron .

RUN ls -al

ENTRYPOINT [ "./sauron" ]
