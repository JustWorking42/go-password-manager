package main

import (
	"context"
	"flag"
	"log"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	var configPath string
	var waitGroup sync.WaitGroup
	flag.StringVar(&configPath, "c", "./server_config.yaml", "path to config file")

	flag.Parse()
	if configPath == "" {
		log.Fatal("config path is required")
	}

	config, err := NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mainContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage, err := mongo.NewMongoStorage(mainContext, config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	auth := auth.NewAuth(config.AuthConfig.Secret)

	creds, err := credintails.Credentials(config.GRPCConfig)
	if err != nil {
		log.Fatal(err)
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
		lis, err := net.Listen("tcp", config.GRPCConfig.Adress())
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("gRPC server is listening on", config.GRPCConfig.Adress())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	stopContext, exit := context.WithTimeout(context.Background(), 5*time.Second)
	storage.Close(stopContext)
	defer exit()

	grpcServer.GracefulStop()
	waitGroup.Wait()
	println("Server exiting")
}
