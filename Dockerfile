# syntax=docker/dockerfile:1
FROM golang:1.17

#ADD . /go/src/manage-order-process
WORKDIR /go/src/github.com/greenbahar/manage-order-process
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy
COPY . ./
RUN go build -o /manage-order-process
EXPOSE 3000
CMD [ "/manage-order-process" ]

#ENTRYPOINT [ "/manage-order-process" ]