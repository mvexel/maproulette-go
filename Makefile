VERSION := $(shell git describe --tags)

test:
	go test ./...

build:
	echo "Building version $(VERSION)"
	go build -ldflags "-X main.Version=$(VERSION)" -o maproulette

release: test build
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: docs
docs:
	godocdown github.com/mvexel/maproulette-go > docs.md
