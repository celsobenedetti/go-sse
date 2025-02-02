# github.com/casey/just
# Default task that lists all available tasks
default: 
	just --list

alias api := dev-api
alias debug := debug-api
alias web := dev-web


# Remove all build artifacts
clean:
	-rm -r ./cmd/**/bin
	-rm -r ./cmd/**/_tmp
	-rm -r ./__debug* 2> /dev/null
	-rm -r ./coverage.out 2> /dev/null
	-find cmd -type f ! -name "*.go*" ! -name "Docker*" -delete 

# ========= Go API =========


# Task to start the API development server using air
dev-api: build-api
	air

# Run Delve debug server
debug-api:
	dlv debug -l 127.0.0.1:8181 --headless ./cmd/api/

# Task to run the API after building it
run-api: build-api
	./cmd/api/bin/main

# Task to build the API
build-api:
	go build -C ./cmd/api/ -o ./bin/main

# Task to clean the API build artifacts
clean-api:
	rm -r ./cmd/api/bin

# Task to run tests for the API
test-api: 
	go test ./... 



# ========= Next client =========

# Task to start the web client development server
dev-web:
	cd web && pnpm dev
