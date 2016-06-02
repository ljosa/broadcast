#!/bin/bash
go run -ldflags "-X github.com/ljosa/broadcast/cmd.version=`git describe --always`" main.go "$@"
