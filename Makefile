GOPATH:=$(CURDIR)/.godeps:$(CURDIR)
export GOPATH

target: dep
	go build -o ./bin/bean ./src/bean/...

debug: dep 
	go build -gcflags "-N -l" -o ./bin/bean ./src/bean/...

dep:
	-mkdir .godeps
	go get github.com/nporsche/np-golang-logging
	go get github.com/nporsche/goyaml

.PHONY: clean test

clean:
	rm -rf bin pkg

test:
	go test protocol -v
	go test util -v
	go test user -v
