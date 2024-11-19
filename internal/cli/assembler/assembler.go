package assembler

import (
	"github.com/spf13/cobra"
	"github.com/yamato0204/grpc-card-server/internal/cli/cmd"
	"github.com/yamato0204/grpc-card-server/internal/cli/cmd/grpc"
	"github.com/yamato0204/grpc-card-server/internal/cli/cmd/grpc/master"
)

// AddSubCommands adds all the subcommands to the rootCmd.
func AddSubCommands(rootCmd *cobra.Command, rootOpts *cmd.RootOption) {
	// rootCmd
	grpcCmd, grpcOpt := grpc.NewGRPCCmd(rootOpts)
	rootCmd.AddCommand(grpcCmd)

	// rootCmd grpc
	grpcMasterCmd, _ := master.NewCmd(grpcOpt)

	grpcCmd.AddCommand(grpcMasterCmd)

}
