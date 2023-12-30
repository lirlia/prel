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
	@echo "Server Running..."
	@go run cmd/prel/main.go

hotrun: db-run # Run the application with hot reload
	@echo "Server Hot Running..."
	@./bin/air

run-e2e: hoverfly-run # Run the application for e2e test
	@echo "Server Running..."
	@HTTP_PROXY=http://localhost:8500 \
		HTTPS_PROXY=http://localhost:8500 \
		NO_PROXY=localhost,127.0.0.1 \
		IS_E2E_MODE=true \
		ADDRESS=127.0.0.1 \
		PROJECT_ID=dummy \
		DB_PASSWORD=password \
		CLIENT_ID=dummy \
		CLIENT_SECRET=dummy \
		PORT=8182 make run

db-run:
	@echo "DB Running..."
	@(cd docker; docker compose up -d)

db-restart:
	@echo "DB Restarting..."
	@(cd docker; docker compose restart)

db-destroy:
	@echo "DB Destroying..."
	@(cd docker; docker compose down)
	@(cd docker; docker compose rm -f)

hoverfly-run:
	@echo "Hoverrun Running..."
	@ps aux | grep hoverfly | grep -vq grep || hoverctl start
	@hoverctl mode simulate
	@hoverctl import ./test/e2e/simulation.json

tidy:
	@echo "Tidying..."
	@go mod tidy

# Generate code
gen: tidy gen-go gen-query # Generate all

gen-go: # Generate go code
	@echo "Go files Generating..."
	@go generate ./...

gen-query: db/sqlc.yaml db/query.sql # Generate query
	@echo "Query Files Generating..."
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
test: test-go test-e2e # Test all

test-go: # Test go
	@echo "Testing..."
	@go test -p 10 -timeout 120s ./...

test-e2e: # Test e2e (need to run server and set proxy to hoverfly)
	@echo "E2E Testing..."
	@npx playwright test -c test/e2e/playwright.config.ts

test-e2e-ui: # Test e2e with ui(for debug)
	@echo "E2E with ui Testing..."
	@npx playwright test --ui -c test/e2e/playwright.config.ts

# Build
build_and_push:
	(cd cmd/prel; KO_DOCKER_REPO=${IMAGE_REGISTRY} ../../bin/ko publish --bare --tags $(IMAGE_TAG) .)

build_and_push_with_timestamp:
	make build_and_push IMAGE_TAG="$(shell date +%s)"
