// Code generated by aip-cli. DO NOT EDIT.
// versions:
// 	protoc        (unknown)
package main

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	v1 "go.einride.tech/aip-cli/cmd/examplectl/einride/example/freight/v1"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "examplectl",
	}
	cmd.AddCommand(v1.NewFreightServiceCommand("freight"))
	return cmd
}

func NewConfig() *aipcli.Config {
	return &aipcli.Config{
		Compiler: aipcli.CompilerConfig{Hosts: map[string]string{}, DefaultHost: "", Root: "examplectl", GoogleCloudIdentityTokens: true, Main: true},
	}
}
