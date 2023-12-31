VERSION := $(shell git describe --tags)

test:
	go test ./...

build:
	echo "Building version $(VERSION)"
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/maproulette

release: test build docs
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: docs
docs:
	go get github.com/robertkrimen/godocdown
	godocdown github.com/mvexel/maproulette-go > docs.md