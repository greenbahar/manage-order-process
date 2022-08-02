#!/bin/bash

docker pull redis
docker run --name m-redis -p 6379:6379 -d redis

cd ..
go mod init manage-order-process
go mod tidy
go run main.go
