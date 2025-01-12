PHONY=test clean all

dev:
	air 

run: build
	./cmd/main/bin/main

build:
	go build -C ./cmd/main/ -o ./bin/main

clean:
	rm -r ./cmd/main/bin
