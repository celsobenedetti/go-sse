# Default task that lists all available tasks
default: 
	just --list


# ========= Go API =========


# Task to start the API development server using air
dev-api: 
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
web:
	cd web && pnpm dev

