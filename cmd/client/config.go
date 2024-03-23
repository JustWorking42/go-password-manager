package main

import (
	"github.com/JustWorking42/go-password-manager/internal/client/grpcclient"
	"github.com/JustWorking42/go-password-manager/internal/client/session"
)

type Config struct {
	GRPCCLientConfig grpcclient.Config `yaml:"grpc"`
	SessionCofig     session.Config    `yaml:"session"`
}
