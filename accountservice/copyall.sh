#!/bin/bash

export GOOS=linux
export CGO_ENABLED=0

# cd accountservice
go get
go build -o accountservice-linux-amd64
echo built `pwd`
# cd ..

export GOOS=darwin

docker build -t ryanyogan/accountservice .

docker service rm accountservice
docker service create \
--name=accountservice \
--replicas=1 \
--publish=6767:6767 \
ryanyogan/accountservice:latest