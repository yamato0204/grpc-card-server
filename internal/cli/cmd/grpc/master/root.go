package master

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yamato0204/grpc-card-server/internal/cli/cmd/grpc"
	pb "github.com/yamato0204/grpc-card-server/pkg/pb/service"
)

type Option struct {
	*grpc.Option
}

func NewCmd(parentOpt *grpc.Option) (*cobra.Command, *Option) {
	opt := &Option{
		Option: parentOpt,
	}

	cmd := &cobra.Command{
		Use:   "master",
		Short: "Call MasterService rpc.",
		Long:  "Call MasterService rpc.",
		Example: strings.Join([]string{
			fmt.Sprintf("%s grpc master <rpc_name> [-d '{}']", opt.RootCmdName),
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("rpc name is required.")
			}

			rpcName := args[0]

			// get data body
			data, err := opt.RequestData()
			if err != nil {
				return err
			}

			grpcClient, err := opt.NewGRPCClient()
			if err != nil {
				return err
			}
			defer grpcClient.Close()

			client := pb.NewMasterServiceClient(grpcClient)

			switch rpcName {
			case "GetAll":
				return handleGetAll(cmd.Context(), opt, client, data)
			case "GetCard":
				return handleGetCard(cmd.Context(), opt, client, data)
			case "GetCharacter":
				return handleGetCharacter(cmd.Context(), opt, client, data)
			}

			return errors.New("rpc name is invalid.")
		},
	}

	return cmd, opt
}

func handleGetAll(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}

	res, err := cli.GetAll(ctx, req)
	if err != nil {
		log.Printf("returns error from server.")
		return nil
	}

	log.Printf("returns success from server.")
	opt.Format(res)
	return nil
}

func handleGetCard(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}

	res, err := cli.GetCard(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}

func handleGetCharacter(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}

	res, err := cli.GetCharacter(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}
