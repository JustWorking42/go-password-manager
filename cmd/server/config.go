package main

import (
	"os"

	"github.com/JustWorking42/go-password-manager/internal/server/auth"
	"github.com/JustWorking42/go-password-manager/internal/server/grpcserver"
	"github.com/JustWorking42/go-password-manager/internal/server/storage"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabaseConfig storage.Config    `yaml:"database"`
	GRPCConfig     grpcserver.Config `yaml:"grpc"`
	AuthConfig     auth.Config       `yaml:"auth"`
}

func NewConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
