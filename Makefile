# Makefile for LibraDemosChain

.PHONY: all build test lint fmt clean

all: build

build:
	cd blockchain && go build ./...

frontend-build:
	cd frontend && npm run build

test:
	cd blockchain && go test ./...
	cd frontend && npm test || echo 'No frontend tests'

lint:
	cd blockchain && golint ./...
	cd frontend && npm run lint || echo 'No frontend linter'

fmt:
	cd blockchain && gofmt -w .
	cd frontend && npm run format || echo 'No frontend formatter'

clean:
	rm -rf blockchain/bin
	echo 'Cleaned build artifacts.'
