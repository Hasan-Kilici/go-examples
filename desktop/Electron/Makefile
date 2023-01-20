.PHONY: all check gocritic build test clean

default: all

all: check test build

test:
	go test -v -coverprofile coverage.out ./...

check:
	pre-commit run --all-files

gocritic:
	gocritic check -enableAll -disable=hugeParam ./...

build:
	go build -ldflags="-w -s"

clean:
	rm -fv ./go-electron-app
