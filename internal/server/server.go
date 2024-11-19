package server

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"syscall"

	masterhandler "github.com/yamato0204/grpc-card-server/internal/handler/master"
	"github.com/yamato0204/grpc-card-server/internal/infra/mysql"
	cardrepository "github.com/yamato0204/grpc-card-server/internal/infra/mysql/repository/card"
	characterrepository "github.com/yamato0204/grpc-card-server/internal/infra/mysql/repository/character"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"

	pb "github.com/yamato0204/grpc-card-server/pkg/pb/service"
)

type Server struct {
	server *grpc.Server
}

type Config struct {
	Port    string `envconfig:"PORT" default:"50051"`
	Profile string `envconfig:"PROFILE" default:"local"`

	// master db
	MysqlMasterAddr     string `envconfig:"MYSQL_MASTER_ADDR" default:"localhost"`
	MysqlMasterProtocol string `envconfig:"MYSQL_MASTER_PROTOCOL" default:"tcp"`
	MysqlMasterUser     string `envconfig:"MYSQL_MASTER_USER" default:"root"`
	MysqlMasterPassword string `envconfig:"MYSQL_MASTER_PASSWORD" default:"root"`
	MysqlMasterDB       string `envconfig:"MYSQL_MASTER_DB" default:"master"`

	// system db
	MysqlShardAddr     string `envconfig:"MYSQL_SHARD_ADDR" default:"localhost"`
	MysqlShardProtocol string `envconfig:"MYSQL_SHARD_PROTOCOL" default:"tcp"`
	MysqlShardUser     string `envconfig:"MYSQL_SHARD_USER" default:"root"`
	MysqlShardPassword string `envconfig:"MYSQL_SHARD_PASSWORD" default:"root"`
	MysqlShardDB       string `envconfig:"MYSQL_SHARD_DB" default:"shard"`
}

func ListenAndServe(conf *Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv, err := newServer(ctx, conf)
	if err != nil {
		return err
	}

	g, gCtx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
		if err != nil {
			return err
		}

		srv.server.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		select {
		case <-gCtx.Done():
		case <-ctx.Done():
			srv.shutdown()
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil

}

func newServer(ctx context.Context, c *Config) (*Server, error) {

	masterDB, err := mysql.NewMasterDB(&mysql.Config{
		Addr:     c.MysqlMasterAddr,
		Protocol: c.MysqlMasterProtocol,
		User:     c.MysqlMasterUser,
		Password: c.MysqlMasterPassword,
		DB:       c.MysqlMasterDB,
	})
	if err != nil {
		return nil, err
	}

	shardDB, err := mysql.NewShardDB(&mysql.Config{
		Addr:     c.MysqlShardAddr,
		Protocol: c.MysqlShardProtocol,
		User:     c.MysqlShardUser,
		Password: c.MysqlShardPassword,
		DB:       c.MysqlShardDB,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			masterDB.Close()
			shardDB.Close()
		}
	}()

	server := grpc.NewServer()
	cardRepository := cardrepository.New(masterDB)
	characterRepository := characterrepository.New(masterDB)

	masterHandler := masterhandler.New(cardRepository, characterRepository)

	pb.RegisterMasterServiceServer(server, masterHandler)

	return &Server{
		server: server,
	}, nil

}

func (srv *Server) shutdown() {
	fmt.Println("shutting down server")
	srv.server.GracefulStop()
}
