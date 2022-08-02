# syntax=docker/dockerfile:1
FROM golang:1.17

#ADD . /go/src/manage-order-process
WORKDIR /go/src/github.com/greenbahar/manage-order-process
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
#RUN go get github.com/greenbahar/manage-order-process
#RUN go install github.com/greenbahar/manage-order-process
RUN go build -o /manage-order-process
#ENTRYPOINT /go/bin/manage-order-process

EXPOSE 3000
CMD [ "/manage-order-process" ]
