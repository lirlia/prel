install: # Install tools
	(cd tools; make install)
	@npm install

clean: clean-test-cache # Clean generated code(without tools)
	@echo "Cleaning..."
	@find . -name "*.gen.*" -type f -delete
	@go clean -testcache

clean-test-cache: # Clean test cache
	@echo "Cleaning..."
	@go clean -testcache

clean-all: clean # Clean all
	@echo "Cleaning..."
	@rm -rf bin

run: db-run # Run the application
	@echo "Running..."
	@go run cmd/prel/main.go

hotrun: db-run # Run the application with hot reload
	@echo "Running..."
	@./bin/air

db-run:
	@echo "Running..."
	@(cd docker; docker compose up -d)

db-restart:
	@echo "Restarting..."
	@(cd docker; docker compose restart)

db-destroy:
	@echo "Destroying..."
	@(cd docker; docker compose down)
	@(cd docker; docker compose rm -f)

hoverfly-run:
	@echo "Running..."
	@ps aux | grep hoverfly | grep -vq grep || hoverctl start
	@hoverctl mode simulate
	@hoverctl import ./test/e2e/simulation.json

tidy:
	@echo "Tidying..."
	@go mod tidy

# Generate code
gen: tidy gen-go gen-query # Generate all

gen-go: # Generate go code
	@echo "Generating..."
	@go generate ./...

gen-query: db/sqlc.yaml db/query.sql # Generate query
	@echo "Generating..."
	@./bin/sqlc generate -f db/sqlc.yaml

# Debug
debug-insert: # Debug insert
	@echo "Debugging..."
	@./scripts/insert-debug-query.sh

# Test / Lint
lint: # Lint
	@echo "Linting..."
	@golangci-lint run ./...

.PHONY: test
test: test-go test-e2e

test-go: # Test go
	@echo "Testing..."
	@go test -p 10 -timeout 120s ./...

test-e2e: hoverfly-run # Test e2e (need to run server and set proxy to hoverfly)
	@echo "Testing..."
	@npx playwright test --retries 3 --fully-parallel --workers 5 --global-timeout 60000

test-e2e-ui: # Test e2e with ui(for debug)
	@echo "Testing..."
	@npx playwright test --ui

# Build
build_and_push:
	(cd cmd/prel; KO_DOCKER_REPO=${IMAGE_REGISTRY} ../../bin/ko publish --bare --tags $(IMAGE_TAG) .)

build_and_push_with_timestamp:
	make build_and_push IMAGE_TAG="$(shell date +%s)"
