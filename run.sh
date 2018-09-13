#!/bin/bash

set -x

pid=`ps aux | grep new-copycat-clean | grep -v grep | awk -F " " '{print $2}'`

kill $pid

git pull

go clean

if [ "$1" == "prod" ];
then
    #GOOS=linux GOARCH=amd64 go build
    go build
    nohup ./go-simple-blog -env="$1"  &
    ps aux | grep go-simple-blog | grep -v grep
else
    go build
   ./go-simple-blog -env="dev"
fi





