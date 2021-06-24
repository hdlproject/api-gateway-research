package main

import (
	"github.com/spf13/cobra"
)

type (
	cmd struct {
		root *cobra.Command
	}
)

func newCmd() *cmd {
	rootCmd := &cobra.Command{
		Use:   "api",
		Short: "API Server",
	}

	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(grpcCmd)

	return &cmd{
		root: rootCmd,
	}
}

func (instance *cmd) execute() {
	if err := instance.root.Execute(); err != nil {
		panic(err)
	}
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP API Server",
	Run: func(cmd *cobra.Command, args []string) {
		httpAPI := newHTTPAPI()
		httpAPI.execute()
	},
}

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "GRPC API Server",
	Run: func(cmd *cobra.Command, args []string) {
		grpcAPI := newGRPCServer()
		grpcAPI.execute()
	},
}
