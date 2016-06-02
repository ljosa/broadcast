#!/bin/bash
if [ "$1" = "linux" ]; then
	export GOOS=linux
	export GOARCH=amd64
fi
go build -ldflags "-X github.com/ljosa/broadcast/cmd.version=`git describe --always`" github.com/ljosa/broadcast
file broadcast
ls -l broadcast
