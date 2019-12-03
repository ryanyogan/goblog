#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

# Account Services
cd accountservice
go get
go build -o accountservice-linux-amd64
echo built `pwd`
cd ..

# Health Checker
cd healthchecker
go get
go build -o healthchecker-linux-amd64
echo built `pwd`
cd ..

cp healthchecker/healthchecker-linux-amd64 accountservice/

export GOOS=darwin

docker build -t ryanyogan/accountservice .
docker push ryanyogan/accountservice:latest

docker service rm accountservice
docker service create \
--name=accountservice \
--replicas=1 \
--publish=6767:6767 \
ryanyogan/accountservice:latest