#!/bin/zsh

cd "$(dirname "$0")"

mkdir -p bin
go build -o bin/wcai main/main.go
