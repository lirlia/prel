//go:build tools
// +build tools

package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/google/ko"
	_ "github.com/ogen-go/ogen"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
	_ "go.uber.org/mock/mockgen"
)
