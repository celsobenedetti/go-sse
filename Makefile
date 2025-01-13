PHONY=test clean all

dev:
	air 

# Run Delve debug server
debug: 
	dlv debug -l 127.0.0.1:8181 --headless ./cmd/api/

run: build
	./cmd/api/bin/main

build:
	go build -C ./cmd/api/ -o ./bin/main

clean:
	rm -r ./cmd/api/bin

