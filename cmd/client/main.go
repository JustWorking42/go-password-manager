package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	"github.com/JustWorking42/go-password-manager/internal/client/commands"
	"github.com/JustWorking42/go-password-manager/internal/client/grpcclient"
	"github.com/JustWorking42/go-password-manager/internal/client/passwordreader"
	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/internal/client/session"
	"github.com/JustWorking42/go-password-manager/internal/common/mapper"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

//go:embed client_config.yaml
var configFile []byte

func main() {
	var rootCmd = rootCmd()
	expandedConfig := os.Expand(string(configFile), mapper.EnvyMapper)

	var config Config
	err := yaml.Unmarshal([]byte(expandedConfig), &config)
	if err != nil {
		log.Fatal(err)
	}
	mainContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := grpcclient.InitAndGetPassGRPCClient(mainContext, config.GRPCCLientConfig)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(client)
	defer client.Close()

	manager, err := session.InitAndGetSessionManager(&config.SessionCofig)
	if err != nil {
		log.Fatal(err)
	}
	session.SetSessionManager(manager)
	defer session.GetSessionManager().Close()

	passwordreader.SetPasswordReader(&passwordreader.TermPasswordReader{})

	addCommands(
		rootCmd,
		commands.RegisterCmd,
		commands.LoginCmd,
		commands.AddPasswordCmd,
		commands.GetPasswordCmd,
		commands.AddCardCmd,
		commands.GetCardCmd,
		commands.AddNoteCmd,
		commands.GetNoteCmd,
		commands.AddBinaryDataCmd,
		commands.GetBinaryDataCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func addCommands(rootCmd *cobra.Command, cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pass",
		Short: "Password manager",
		Long:  "Password manager",
	}
}
