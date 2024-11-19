package output

import (
	"errors"
	"io"

	ojson "github.com/yamato0204/grpc-card-server/internal/cli/output/json"
)

type Output interface {
	// Log
	Log(msg string)
	// Logf
	Logf(format string, args ...interface{})
	// Format
	Format(data interface{})
}

// NewOutputWithFormatter
func NewOutputWithFormatter(formatter string, writer io.Writer) (Output, error) {
	switch formatter {
	case "json":
		return ojson.NewOutput(writer), nil

	default:
		return nil, errors.New("invalid formatter")
	}
}
