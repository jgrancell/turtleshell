#!/bin/sh

BINARY=$1
WORKDIR=$2

cd ${WORKDIR}

GOOS=linux GOARCH=amd64 go build -o packaging/binaries/${BINARY}-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o packaging/binaries/${BINARY}-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o packaging/binaries/${BINARY}-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o packaging/binaries/${BINARY}-windows-amd64

cd packaging/binaries

tar -czvf ${BINARY}-linux-amd64.tar.gz ${BINARY}-linux-amd64
tar -czvf ${BINARY}-darwin-amd64.tar.gz ${BINARY}-darwin-amd64
tar -czvf ${BINARY}-darwin-arm64.tar.gz ${BINARY}-darwin-arm64
tar -czvf ${BINARY}-windows-amd64.tar.gz ${BINARY}-windows-amd64