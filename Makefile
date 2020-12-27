SHELL=/bin/bash

start:
	go run main.go

dev:
	air

build:
	go build main.go

test:
	go test main_test.go

clean:
	rm main