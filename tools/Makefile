install: go.mod tools.go
	@GOBIN=$(shell pwd)/../bin go install github.com/air-verse/air
	@GOBIN=$(shell pwd)/../bin go install github.com/sqlc-dev/sqlc/cmd/sqlc
	@GOBIN=$(shell pwd)/../bin go install github.com/ogen-go/ogen/cmd/ogen
	@GOBIN=$(shell pwd)/../bin go install github.com/google/ko
	@GOBIN=$(shell pwd)/../bin go install go.uber.org/mock/mockgen
