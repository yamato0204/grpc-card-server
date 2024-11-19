package cli

import (
	"fmt"
	"os"

	"github.com/yamato0204/grpc-card-server/internal/cli/assembler"
	"github.com/yamato0204/grpc-card-server/internal/cli/cmd"
)

func main() {
	rootCmd, opt := cmd.NewCmdRoot(os.Stdout, os.Stderr)
	assembler.AddSubCommands(rootCmd, opt)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("command execution error: %v\n", err)
		os.Exit(1)
	}
}
