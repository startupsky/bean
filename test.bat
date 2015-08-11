SET GOPATH=%cd%/.godeps;%cd%
go test protocol -v
go test util -v
go test user -v