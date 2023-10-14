#!/bin/sh

GOOS=windows GOARCH=amd64 go build cmd/client/client.go

echo "Press enter to continue";
read name;
