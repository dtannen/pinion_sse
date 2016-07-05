#!/bin/bash
export GOPATH=`pwd`
echo "[-] Get Packages"
go get github.com/dtannen/sseserver
echo "[-] Go Build"
go build -o pinion_sse
