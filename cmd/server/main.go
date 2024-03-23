package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/JustWorking42/go-password-manager/internal/server/auth"
	"github.com/JustWorking42/go-password-manager/internal/server/credintails"
	"github.com/JustWorking42/go-password-manager/internal/server/grpcserver"
	"github.com/JustWorking42/go-password-manager/internal/server/storage/mongo"
	"github.com/JustWorking42/go-password-manager/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	var configPath string
	var waitGroup sync.WaitGroup
	flag.StringVar(&configPath, "c", "./server_config.yaml", "path to config file")

	flag.Parse()
	if configPath == "" {
		logrus.Fatal("no config file provided")
	}

	config, err := NewConfig(configPath)
	if err != nil {
		logrus.WithError(err).Fatal("failed to load config")
	}

	mainContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage, err := mongo.NewMongoStorage(mainContext, config.DatabaseConfig)
	if err != nil {
		logrus.WithError(err).Fatal("failed to initialize storage")
	}

	auth := auth.NewAuth(config.AuthConfig.Secret)

	creds, err := credintails.Credentials(config.GRPCConfig)
	if err != nil {
		logrus.WithError(err).Fatal("failed to initialize credentials")
	}

	interceptors := grpc.ChainUnaryInterceptor(
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			if info.FullMethod == proto.PassManager_Login_FullMethodName || info.FullMethod == proto.PassManager_Register_FullMethodName {
				return handler(ctx, req)
			}
			return auth.OnlyAuthorizedInterceptor(ctx, req, info, handler)
		},
	)
	grpcServer := grpc.NewServer(grpc.Creds(creds), interceptors)

	proto.RegisterPassManagerServer(grpcServer, grpcserver.NewPassGRPCServer(storage, auth))
	reflection.Register(grpcServer)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GRPCConfig.Host, config.GRPCConfig.Port))
		if err != nil {
			logrus.WithError(err).Fatal("failed to listen")
		}
		logrus.Infof("gRPC server is listening on %s:%s", config.GRPCConfig.Host, config.GRPCConfig.Port)
		if err := grpcServer.Serve(lis); err != nil {
			logrus.WithError(err).Fatal("failed to serve")
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	stopContext, exit := context.WithTimeout(context.Background(), 5*time.Second)
	err = storage.Close(stopContext)
	if err != nil {
		logrus.WithError(err).Error("failed to close storage")
	}
	defer exit()

	grpcServer.GracefulStop()
	waitGroup.Wait()
	logrus.Info("gRPC server stopped")
}
