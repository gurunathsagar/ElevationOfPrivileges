#!/bin/sh

# Installing go
sudo apt-get install python-software-properties
sudo add-apt-repository ppa:duh/golang
sudo apt-get update
sudo apt-get install golang
#go version
#go version go1.1.1 linux/amd64

export GOROOT=/usr/lib/go
export GOBIN=/usr/bin/go

# Install Mux. This should be run in workspace/src
go get -u github.com/gorilla/mux