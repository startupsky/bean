@echo off
SET GOPATH=%cd%/.godeps;%cd%
rmdir /s /q .godeps
mkdir .godeps
go get github.com/op/go-logging
go get github.com/nporsche/goyaml