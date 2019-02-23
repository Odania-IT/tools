.PHONY: build test clean

VERSION=$(shell git describe --tags)

build: codepipeline-status_amd64 codepipeline-status_alpine codepipeline-status_darwin codepipeline-status.exe

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

clean:
	rm -f .get-deps
	rm -f *_amd64 *_darwin *.exe

test: .get-deps *.go
	go test -v ./...

codepipeline-status_amd64: .get-deps *.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@ *.go

codepipeline-status_alpine: .get-deps *.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ *.go

codepipeline-status_darwin: .get-deps *.go
	GOOS=darwin go build -o $@ *.go

codepipeline-status.exe: .get-deps *.go
	GOOS=windows GOARCH=amd64 go build -o $@ *.go
